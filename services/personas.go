package services

import (
	"github.com/maxilovera/go-crud-example/dto"
)

var personas []*dto.Persona

func ObtenerPersonas() []*dto.Persona {
	return personas
}

func ObtenerPersonaPorId(id int) *dto.Persona {
	for _, persona := range personas {
		if persona.ID == id {
			return persona
		}
	}

	return nil
}

func CrearPersona(nuevaPersona dto.Persona) *dto.Persona {

	nuevaPersona.ID = len(personas) + 1

	personas = append(personas, &nuevaPersona)

	return &nuevaPersona
}

func ActualizarPersona(personaActualizada dto.Persona) (string) {
	//Tengo que buscar la persona por ID, y modificarle los datos con los de la nueva
	personaPorActualizar := ObtenerPersonaPorId(personaActualizada.ID)

	if (personaPorActualizar == nil) {
		return "persona no existe"
	}

	personaPorActualizar.Nombre = personaActualizada.Nombre
	personaPorActualizar.Edad = personaActualizada.Edad

	return ""
}

func EliminarPersona(id int) string{
	// Index of the element to be deleted
	indexToDelete := 2

	if indexToDelete >= 0 && indexToDelete < len(personas) {
		// Creating a new slice excluding the element to be deleted
		modifiedSlice := append(personas[:indexToDelete], personas[indexToDelete+1:]...)
		personas = modifiedSlice
		return ""
	}

	return "Error"
}