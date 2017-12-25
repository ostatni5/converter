// Package serialize provide Serialize function to convert model -> SHIELD input files.
package serialize

import (
	"github.com/yaptide/converter/log"
	"github.com/yaptide/converter/setup"
	"github.com/yaptide/converter/shield"
	"github.com/yaptide/converter/shield/setup/serialize/data"
)

// Result store result of Serialize() function.
type Result struct {
	// SimulationContext is necessary to bind shield simulation results to model.
	SimulationContext shield.SimulationContext

	//  Files map where: key is file name, value is file content.
	Files map[string]string
}

// Serialize simulation setup.Setup to shield input files.
func Serialize(setup setup.Setup) (Result, error) {
	data, simulationContext, err := data.Convert(setup)
	if err != nil {
		return Result{}, err
	}

	return Result{
		SimulationContext: simulationContext,
		Files:             serializeData(data),
	}, nil
}

func serializeData(data data.Data) map[string]string {
	log.Debug("[Serializer] data %+v", data)
	files := map[string]string{}

	for fileName, serializeFunc := range map[string]func() string{
		materialsDatFile: func() string { return serializeMat(data.Materials) },
		geometryDatFile:  func() string { return serializeGeo(data.Geometry) },
		detectorsDatFile: func() string { return serializeDetect(data.Detectors) },
		beamDatFile:      func() string { return serializeBeam(data.Beam, data.Options) },
	} {
		serializeOutput := serializeFunc()
		files[fileName] = serializeOutput
	}

	log.Debug("Files:\n")
	for filename, content := range files {
		log.Debug("%s:\n%s\n\n", filename, content)
	}

	return files
}
