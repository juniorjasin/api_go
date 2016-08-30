package sort

import(
	"errors"
)


func BubbleSort(numeros []int){
	if numeros == nil {
	//TODO: retornar un error
	return errors.New("No se puede procesar un array vacio")
	}
	
	Listo := true
	for Listo {
		Listo = false;
		for i := 0; i < (len(numeros) - 1); i++ {
			if numeros[i] < numeros[i+1] {
				numeros[i], numeros[i+1] = numeros[i+1], numeros[i]
				Listo = true
			}
		}
	}

	return nil
}