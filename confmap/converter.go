// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package confmap // import "go.opentelemetry.io/collector/confmap"

import (
	"context"
)

// ConverterSettings are the settings to initialize a Converter.
type ConverterSettings struct{}

// ConverterFactory defines a factory that can be used to instantiate
// new instances of a Converter.
type ConverterFactory = moduleFactory[Converter, ConverterSettings]

// CreateConverterFunc is a function that creates a Converter instance.
type CreateConverterFunc = createConfmapFunc[Converter, ConverterSettings]

// NewConverterFactory can be used to create a ConverterFactory.
func NewConverterFactory(f CreateConverterFunc) ConverterFactory {
	return newConfmapModuleFactory(f)
}

// Converter is a converter interface for the confmap.Conf that allows distributions
// (in the future components as well) to build backwards compatible config converters.
type Converter interface {
	// Convert applies the conversion logic to the given "conf".
	Convert(ctx context.Context, conf *Conf) error
}
