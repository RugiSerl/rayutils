package graphic

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

// Object representing a shader program, make shader usage easier
type Shader struct {
	uniformLocations map[string]int32
	Shader           rl.Shader
}

func InitShader(vertexShaderPath, fragmentShaderPath string) *Shader {
	s := new(Shader)
	s.Shader = rl.LoadShader(vertexShaderPath, fragmentShaderPath)
	s.uniformLocations = make(map[string]int32)
	return s

}

func (s *Shader) SetValueFromUniformName(uniformName string, value any, uniformType rl.ShaderUniformDataType) {
	if location, ok := s.uniformLocations[uniformName]; ok {
		s.setValue(location, value, uniformType)

	} else {
		s.uniformLocations[uniformName] = rl.GetShaderLocation(s.Shader, uniformName)
		s.setValue(s.uniformLocations[uniformName], value, uniformType)

	}

}
func (s *Shader) setValue(uniformLocation int32, value any, uniformType rl.ShaderUniformDataType) {
	switch value.(type) {
	case float32:
		rl.SetShaderValue(s.Shader, uniformLocation, []float32{value.(float32)}, uniformType)
	case []float32:
		rl.SetShaderValue(s.Shader, uniformLocation, value.([]float32), uniformType)
	}

}

// Begin custom shader drawing.
func (s *Shader) Begin() {
	rl.BeginShaderMode(s.Shader)
}

// End custom shader drawing.
func (s *Shader) End() {
	rl.EndShaderMode()
}
