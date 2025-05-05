package objectresource

import (
	"fmt"

	"gopkg.in/yaml.v3"
)

type Validation struct {
	Rules   []ValidationRuleInterface
	decoder *yaml.Decoder
}

func NewValidation(rules ...ValidationRuleInterface) Validation {
	return Validation{
		Rules: rules,
	}
}

func (s *Validation) UnmarshalYAML(node *yaml.Node) error {

	if node.Kind == yaml.SequenceNode {
		for _, content := range node.Content {
			if content.Kind == yaml.MappingNode {
				for i := 0; i < len(content.Content); i += 2 {
					keyNode := content.Content[i]
					valueNode := content.Content[i+1]

					key, err := getKeyFromNode(keyNode)
					if err != nil {
						return err
					}
					value, err := getValueFromNode(valueNode)
					if err != nil {
						return err
					}

					if key == "Type" && value == "Regex" || key == "Regex" {
						rule, err := getRuleFromNode[ValidationRegexRule](content)
						if err != nil {
							return err
						}
						s.Rules = append(s.Rules, rule)
						break
					} else if key == "Type" && value == "Length" || key == "Length" {
						rule, err := getRuleFromNode[ValidationLengthRule](content)
						if err != nil {
							return err
						}
						s.Rules = append(s.Rules, rule)
						break
					}
				}
			}
		}
	}

	return nil
}
func getKeyFromNode(node *yaml.Node) (string, error) {
	if node.Kind != yaml.ScalarNode {
		return "", fmt.Errorf("key is not scalar node")
	}
	return node.Value, nil
}

func getValueFromNode(node *yaml.Node) (interface{}, error) {
	var value interface{}
	err := yaml.Unmarshal([]byte(node.Value), &value)
	if err != nil {
		return nil, err
	}
	return value, nil
}

func getRuleFromNode[T ValidationRuleInterface](node *yaml.Node) (T, error) {
	var value T
	err := node.Decode(&value)
	if err != nil {
		return value, err
	}
	return value, nil
}
