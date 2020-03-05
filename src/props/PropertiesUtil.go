package props

import (
	"bufio"
	"os"
	"strings"
	"log"
)

type AppConfigProperties map[string]string

//
// Reads the given file into a map
//
func ReadPropertiesFile(filename string) (AppConfigProperties, error) {

	config := AppConfigProperties{}

	if len(filename) == 0 {
		return config, nil
	}
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	// For each line..
	for scanner.Scan() {

		// Get the text..
		line := scanner.Text()

		// If there is an equal in there...
		if equal := strings.Index(line, "="); equal >= 0 {

			// We slice the line array from 0 to the = sign, and trim it.
			// slice[low : high]
			// If the key has a length > 0 after trimming then
			if key := strings.TrimSpace(line[:equal]); len(key) > 0 {
				value := ""

				// If there are characters after the equals sign
				if len(line) > equal {
					// Slice the line from the equal to the end.
					value = strings.TrimSpace(line[equal+1:])
				}

				// Add do the map
				config[key] = value
			}
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
		return nil, err
	}

	return config, nil
}
