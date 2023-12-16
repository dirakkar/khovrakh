/*
 * Copyright © 2023 Dirakkar
 *
 * This source code is licensed under the MIT license found in the LICENSE
 * file in the root directory of this source tree.
 */

package config

import (
	"os"
	"path/filepath"

	"github.com/mitchellh/mapstructure"

	"github.com/dirakkar/khovrakh/config/adapter"
	"github.com/dirakkar/khovrakh/config/hook"
)

func New[T any](filename string, adapters ...adapter.Adapter) (*T, error) {
	var result T

	decoderConfig := mapstructure.DecoderConfig{
		TagName: "cfg",
		Result:  &result,
		DecodeHook: mapstructure.ComposeDecodeHookFunc(
			mapstructure.StringToTimeDurationHookFunc(),
			hook.StringToSize(),
		),
	}

	decoder, err := mapstructure.NewDecoder(&decoderConfig)
	if err != nil {
		return nil, err
	}

	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	adapters = append(adapters, adapter.JSON{})

	output, err := Unmarshal(UnmarshalInput{
		Type:     filepath.Ext(filename),
		Data:     data,
		Adapters: adapters,
	})
	if err != nil {
		return nil, err
	}

	if err := decoder.Decode(output); err != nil {
		return nil, err
	}

	return &result, nil
}

type UnmarshalInput struct {
	Type     string
	Data     []byte
	Adapters []adapter.Adapter
}

func Unmarshal(input UnmarshalInput) (map[string]any, error) {
	var result map[string]any

	for _, adapter := range input.Adapters {
		if !adapter.Peek(input.Type) {
			continue
		}

		if err := adapter.Unmarshal(input.Data, &result); err != nil {
			return nil, err
		}

		break
	}

	return result, nil
}
