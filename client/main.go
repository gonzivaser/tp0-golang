package main

import (
	"client/globals"
	"client/utils"
	"log"
)

func main() {
	utils.ConfigurarLogger()

	// loggear "Hola soy un log" usando la biblioteca log
	globals.ClientConfig = utils.IniciarConfiguracion("config.json") // ESTO ABRE EL ARCHIVO
	log.Println("Soy un log")

	// validar que la config este cargada correctamente
	if globals.ClientConfig == nil {
		log.Fatalf("No se pudo cargar la configuración") // FATALF CORTA EJECUCION
	}

	// loggeamos el valor de la config
	log.Printf("Informacion cargada correctamente %s", globals.ClientConfig.Mensaje)

	// ADVERTENCIA: Antes de continuar, tenemos que asegurarnos que el servidor esté corriendo para poder conectarnos a él
	// EN UTILS.GO DE SERVER, BORRE LA LINEA DONDE LLAMA A PANIC, PARA QUE EL SERVIDOR ARRANQUE A CORRER

	// enviar un mensaje al servidor con el valor de la config
	utils.EnviarMensaje(globals.ClientConfig.Ip, globals.ClientConfig.Puerto, globals.ClientConfig.Mensaje)

	// leer de la consola el mensaje
	// utils.LeerConsola()
	mensajes := utils.LeerConsola()

	// generamos un paquete y lo enviamos al servidor
	// utils.GenerarYEnviarPaquete()
	utils.GenerarYEnviarPaquete(mensajes)

}
