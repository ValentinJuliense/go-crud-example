package services

import (
	"github.com/go-crud-example/dto"
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

func ActualizarPersona(id int, nuevaPersona *dto.Persona) *dto.Persona {
	persona := ObtenerPersonaPorId(id)
	if persona != nil {
		persona.Nombre = nuevaPersona.Nombre
		persona.Edad = nuevaPersona.Edad
		return persona
	}

	return nil
}

func EliminarPersona(id int) []*dto.Persona {
	persona := ObtenerPersonaPorId(id)
	indice := -1
	if persona != nil {
		for i, obj := range personas {
			if persona.ID == obj.ID {
				indice = i
				break
			}
		}
		//personas[indice] = personas[len(personas)-1] //mueve el elemento que quiero eliminar a la ultima posicion
		//personas = personas[:len(personas)-1]        //no copia la lista, solo corta el ultimo elemento

		personas = append(personas[:indice], personas[indice+1:]...) // problema: listas grandes

		return personas
	}
	return nil
}
