package main

type Embarazada struct {
	Id                     int    `json:"id"`
	NoExpediente           string `json:"noExpediente"`
	Nombre                 string `json:"nombre"`
	Curp                   string `json:"curp"`
	Direccion              string `json:"domicilioReferencia"`
	Telefono               int64  `json:"telefono,string,omitempty"`
	FechaNacimiento        string `json:"FechaNacimiento,omitempty"`
	Gestas                 string `json:"gestas"`
	Paras                  string `json:"paras"`
	Abortos                string `json:"abortos"`
	Cesareas               string `json:"cesareas"`
	DondeMigro             string `json:"dondeEmigro"`
	ConsultaPregestacional string `json:"ConsultaPregestacional,omitempty"`
}

func insertarEmbarazada(c Embarazada) (e error) {
	db, err := obtenerBaseDeDatos()
	if err != nil {
		return err
	}
	defer db.Close()

	// Preparamos para prevenir inyecciones SQL
	sentenciaPreparada, err := db.Prepare("INSERT INTO mujer (No_Expediente, Nombre,FechaNacimiento, curp, Telefono, Domicilio_Referencia, Gestas, Paras, Abortos, Cesareas, Donde_Emigro, Consulta_RiesgoPreg) VALUES(?,?,?,?,?,?,?,?,?,?,?,?)")
	if err != nil {
		return err
	}
	defer sentenciaPreparada.Close()
	// Ejecutar sentencia, un valor por cada '?'
	_, err = sentenciaPreparada.Exec(c.NoExpediente, c.Nombre, c.FechaNacimiento, c.Curp, c.Telefono, c.Direccion, c.Gestas, c.Paras, c.Abortos, c.Cesareas, c.DondeMigro, c.ConsultaPregestacional)
	if err != nil {
		return err
	}
	return nil
}

func establecerFechaSalida(IdVehiculo int64, FechaSalida string) error {
	bd, err := obtenerBaseDeDatos()
	if err != nil {
		return err
	}
	_, err = bd.Exec("UPDATE vehiculos SET fecha_salida = ? WHERE id = ?", FechaSalida, IdVehiculo)
	if err != nil {
		return err
	}
	return nil
}

func obtenerVehiculos() ([]Embarazada, error) {
	vehiculos := []Embarazada{}
	bd, err := obtenerBaseDeDatos()

	if err != nil {
		return vehiculos, err
	}

	defer bd.Close()
	filas, err := bd.Query(`SELECT id_mujer,No_Expediente, nombre,curp,telefono,  COALESCE(FechaNacimiento, ''), domicilio_referencia, gestas, paras, abortos, cesareas, dondeEmigro, COALESCE(ConsultaPregestacional, '') FROM mujer`)
	if err != nil {
		return vehiculos, err
	}
	defer filas.Close()
	var vehiculo Embarazada
	for filas.Next() {
		err := filas.Scan(&vehiculo.Id, &vehiculo.NoExpediente, &vehiculo.Nombre, &vehiculo.Curp, &vehiculo.Telefono, &vehiculo.FechaNacimiento, &vehiculo.Direccion, &vehiculo.Gestas, &vehiculo.Paras, &vehiculo.Abortos, &vehiculo.Cesareas, &vehiculo.DondeMigro, &vehiculo.ConsultaPregestacional)

		if err != nil {
			return vehiculos, err
		}
		vehiculos = append(vehiculos, vehiculo)

	}
	return vehiculos, nil
}
