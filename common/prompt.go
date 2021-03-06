package common

import "github.com/manifoldco/promptui"

type Field struct {
	Name   string
	Prompt promptui.Prompt
}

func ReadMap(mapping []Field) (map[string]string, error) {
	result := make(map[string]string)

	for _, value := range mapping {
		ret, err := value.Prompt.Run()
		if err != nil {
			return nil, err
		}

		result[value.Name] = ret
	}

	return result, nil
}
