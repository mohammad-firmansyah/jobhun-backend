package api

import (
	"database/sql"
	"fmt"
	"strconv"

	"github.com/go-playground/validator"
	"github.com/gofiber/fiber/v2"
	"github.com/mohammad-firmansyah/jobhun-backend/database"
)

var (
	mahasiswas = map[string]database.Mahasiswa{}
	mysql      *sql.DB
)

func createJurusan(c *fiber.Ctx) error {
	jurusan := new(database.Jurusan)

	err := c.BodyParser(&jurusan)

	if err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
			"errors": err.Error(),
		})

	}

	validate := validator.New()
	err = validate.Struct(jurusan)
	if err != nil {
		c.Status(404).JSON(&fiber.Map{
			"error": err.Error(),
		})
		return nil
	}

	data, err := mysql.Exec(`INSERT INTO jurusan
      VALUES ('','` + jurusan.NamaJurusan + `')`)

	if err != nil {
		panic(err)
	}

	id_jurusan, err := data.LastInsertId()

	jurusan.ID = strconv.FormatInt(id_jurusan, 10)
	c.Status(200).JSON(&fiber.Map{
		"jurusan": jurusan,
	})

	return nil
}

func readJurusans(c *fiber.Ctx) error {

	rows, err := mysql.Query("SELECT * FROM jurusan")
	var listJurusan []database.Jurusan
	if err != nil {
		panic(err)
	}

	for rows.Next() {
		var jurusan database.Jurusan

		err := rows.Scan(&jurusan.ID, &jurusan.NamaJurusan)
		if err != nil {
			panic(err)
		}

		listJurusan = append(listJurusan, jurusan)
	}

	c.Status(200).JSON(&fiber.Map{
		"jurusan": listJurusan,
	})
	return nil
}

func createHobi(c *fiber.Ctx) error {
	hobi := new(database.Hobi)

	err := c.BodyParser(&hobi)

	if err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
			"errors": err.Error(),
		})

	}

	validate := validator.New()
	err = validate.Struct(hobi)
	if err != nil {
		c.Status(404).JSON(&fiber.Map{
			"error": err.Error(),
		})
		return nil
	}

	data, err := mysql.Exec(`INSERT INTO hobi
      VALUES ('','` + hobi.NamaHobi + `')`)

	if err != nil {
		panic(err)
	}

	id_hobi, err := data.LastInsertId()

	hobi.ID = strconv.FormatInt(id_hobi, 10)

	c.Status(200).JSON(&fiber.Map{
		"hobi": hobi,
	})

	return nil
}

func readHobis(c *fiber.Ctx) error {

	rows, err := mysql.Query("SELECT * FROM hobi")
	var listHobi []database.Hobi
	if err != nil {
		panic(err)
	}

	for rows.Next() {
		var hobi database.Hobi

		err := rows.Scan(&hobi.ID, &hobi.NamaHobi)
		if err != nil {
			panic(err)
		}

		listHobi = append(listHobi, hobi)
	}

	c.Status(200).JSON(&fiber.Map{
		"hobi": listHobi,
	})
	return nil
}

func createMahasiswa(c *fiber.Ctx) error {
	mahasiswa := new(database.Mahasiswa)

	err := c.BodyParser(&mahasiswa)

	if err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
			"errors": err.Error(),
		})

	}

	validate := validator.New()
	err = validate.Struct(mahasiswa)
	if err != nil {
		c.Status(404).JSON(&fiber.Map{
			"error": err.Error(),
		})
		return nil
	}

	row := mysql.QueryRow("SELECT * FROM jurusan WHERE nama_jurusan = '" + mahasiswa.Jurusan + "'")

	var jurusan database.Jurusan

	err = row.Scan(&jurusan.ID, &jurusan.NamaJurusan)
	if jurusan.NamaJurusan != mahasiswa.Jurusan {
		c.Status(404).JSON(&fiber.Map{
			"error": "jurusan not found",
		})
		return nil
	}

	if err != nil {

		panic(err.Error())
	}

	data, err := mysql.Exec(`INSERT INTO mahasiswa
      VALUES ('','` + mahasiswa.Nama + `' , '` + strconv.Itoa(mahasiswa.Usia) + `','` + strconv.Itoa(mahasiswa.Gender) + `' , '` + mahasiswa.TglRegistrasi + `', '` + jurusan.ID + `')`)

	if err != nil {
		panic(err)
	}

	for _, v := range mahasiswa.Hobi {
		row := mysql.QueryRow("SELECT * FROM hobi WHERE nama_hobi = '" + v + "'")

		var hobi database.Hobi
		err := row.Scan(&hobi.ID, &hobi.NamaHobi)
		if hobi.NamaHobi != v {
			c.Status(404).JSON(&fiber.Map{
				"error": "hobi not found",
			})
			return nil
		}

		if err != nil {
			panic(err)
		}

		id_mahasiswa, err := data.LastInsertId()
		_, err = mysql.Exec(`INSERT INTO mahasiswa_hobi
      VALUES ('','` + strconv.FormatInt(id_mahasiswa, 10) + `' ,'` + hobi.ID + `' )`)
		if err != nil {
			panic(err)
		}

	}

	c.Status(200).JSON(&fiber.Map{
		"mahasiswa": mahasiswa,
	})

	return nil
}

func readMahasiswa(c *fiber.Ctx) error {
	id := c.Params("id")
	row := mysql.QueryRow("SELECT * FROM mahasiswa WHERE id = " + id)

	var mahasiswa database.Mahasiswa

	err := row.Scan(&mahasiswa.ID, &mahasiswa.Nama, &mahasiswa.Usia, &mahasiswa.Gender, &mahasiswa.TglRegistrasi, &mahasiswa.Jurusan)
	if mahasiswa.ID != id {
		c.Status(404).JSON(&fiber.Map{
			"error": "mahasiswa not found ",
		})
		return nil
	}

	if err != nil {

		panic(err.Error())
	}

	// update pivot table
	row = mysql.QueryRow("SELECT * FROM jurusan WHERE id = " + mahasiswa.Jurusan)

	var jurusan database.Jurusan

	err = row.Scan(&jurusan.ID, &jurusan.NamaJurusan)
	if jurusan.ID != mahasiswa.Jurusan {
		c.Status(404).JSON(&fiber.Map{
			"error": "jurusan not found",
		})
		return nil
	}

	if err != nil {

		panic(err.Error())
	}

	rows, err := mysql.Query("SELECT * FROM mahasiswa_hobi WHERE id_mahasiswa = " + id)
	var listHobi []string

	if err != nil {
		panic(err)
	}

	for rows.Next() {
		var mahasiswaHobi database.MahasiswaHobi

		err := rows.Scan(&mahasiswaHobi.ID, &mahasiswaHobi.IdMahasiswa, &mahasiswaHobi.IdHobi)
		if err != nil {
			panic(err)
		}

		row := mysql.QueryRow("SELECT * FROM hobi WHERE id = " + mahasiswaHobi.IdHobi)

		var hobi database.Hobi

		err = row.Scan(&hobi.ID, &hobi.NamaHobi)
		if hobi.ID != mahasiswaHobi.IdHobi {
			c.Status(404).JSON(&fiber.Map{
				"error": "hobi not found",
			})
			return nil
		}

		if err != nil {

			panic(err.Error())
		}

		listHobi = append(listHobi, hobi.NamaHobi)

	}

	mahasiswa.Hobi = listHobi
	mahasiswa.Jurusan = jurusan.NamaJurusan

	c.Status(200).JSON(&fiber.Map{
		"mahasiswa": mahasiswa,
	})
	return nil

}

func readMahasiswas(c *fiber.Ctx) error {
	rows, err := mysql.Query("SELECT * FROM mahasiswa")
	var listMahasiswa []database.Mahasiswa
	if err != nil {
		panic(err)
	}

	for rows.Next() {
		var mahasiswa database.Mahasiswa

		err := rows.Scan(&mahasiswa.ID, &mahasiswa.Nama, &mahasiswa.Usia, &mahasiswa.Gender, &mahasiswa.TglRegistrasi, &mahasiswa.Jurusan)
		if err != nil {
			panic(err)
		}

		// update pivot table

		fmt.Print("jurusan id : " + mahasiswa.Jurusan)
		row := mysql.QueryRow("SELECT * FROM jurusan WHERE id = " + mahasiswa.Jurusan)

		var jurusan database.Jurusan

		err = row.Scan(&jurusan.ID, &jurusan.NamaJurusan)
		if jurusan.ID != mahasiswa.Jurusan {
			c.Status(404).JSON(&fiber.Map{
				"error": "jurusan not found",
			})
			return nil
		}

		if err != nil {

			panic(err.Error())
		}

		rows, err := mysql.Query("SELECT * FROM mahasiswa_hobi WHERE id_mahasiswa = " + mahasiswa.ID)

		if err != nil {
			panic(err)
		}
		var listHobi []string
		for rows.Next() {
			var mahasiswaHobi database.MahasiswaHobi

			err := rows.Scan(&mahasiswaHobi.ID, &mahasiswaHobi.IdMahasiswa, &mahasiswaHobi.IdHobi)
			if err != nil {
				panic(err)
			}

			row := mysql.QueryRow("SELECT * FROM hobi WHERE id = " + mahasiswaHobi.IdHobi)

			var hobi database.Hobi

			err = row.Scan(&hobi.ID, &hobi.NamaHobi)
			if hobi.ID != mahasiswaHobi.IdHobi {
				c.Status(404).JSON(&fiber.Map{
					"error": "hobi not found",
				})
				return nil
			}

			if err != nil {

				panic(err.Error())
			}

			listHobi = append(listHobi, hobi.NamaHobi)

		}
		mahasiswa.Hobi = listHobi
		mahasiswa.Jurusan = jurusan.NamaJurusan

		listMahasiswa = append(listMahasiswa, mahasiswa)
	}

	c.Status(200).JSON(&fiber.Map{
		"mahasiswa": listMahasiswa,
	})
	return nil
}

func updateMahasiswa(c *fiber.Ctx) error {
	updateMahasiswa := new(database.Mahasiswa)

	err := c.BodyParser(updateMahasiswa)

	if err != nil {
		c.Status(500).JSON(&fiber.Map{
			"error": err.Error(),
		})
		return err
	}

	fmt.Println(updateMahasiswa.Nama)

	id := c.Params("id")
	row := mysql.QueryRow("SELECT * FROM mahasiswa WHERE id = '" + id + "'")

	var mahasiswa database.Mahasiswa

	err = row.Scan(&mahasiswa.ID, &mahasiswa.Nama, &mahasiswa.Usia, &mahasiswa.Gender, &mahasiswa.TglRegistrasi, &mahasiswa.Jurusan)
	if mahasiswa.ID != id {
		c.Status(404).JSON(&fiber.Map{
			"error": "mahasiswa not found",
		})
		return nil
	}

	if err != nil {

		panic(err.Error())
	}

	row = mysql.QueryRow("SELECT * FROM jurusan WHERE id = " + mahasiswa.Jurusan)

	var jurusan database.Jurusan

	err = row.Scan(&jurusan.ID, &jurusan.NamaJurusan)
	if jurusan.ID != mahasiswa.Jurusan {
		c.Status(404).JSON(&fiber.Map{
			"error": "jurusan not found",
		})
		return nil
	}

	if err != nil {

		panic(err.Error())
	}

	// update table mahasiswa
	_, err = mysql.Exec(`UPDATE mahasiswa
      SET nama = '` + updateMahasiswa.Nama + `',usia = ` + strconv.Itoa(updateMahasiswa.Usia) + `,gender = ` + strconv.Itoa(updateMahasiswa.Gender) + `,tgl_registrasi = '` + updateMahasiswa.TglRegistrasi + `',id_jurusan = ` + jurusan.ID + ` WHERE id = ` + id)
	if err != nil {
		panic(err)
	}

	// update pivot table
	_, err = mysql.Exec(`DELETE FROM mahasiswa_hobi WHERE id_mahasiswa = ` + id)
	if err != nil {
		panic(err)
	}

	for _, v := range updateMahasiswa.Hobi {
		row := mysql.QueryRow("SELECT * FROM hobi WHERE nama_hobi = '" + v + "'")

		var hobi database.Hobi
		err := row.Scan(&hobi.ID, &hobi.NamaHobi)
		if hobi.NamaHobi != v {
			c.Status(404).JSON(&fiber.Map{
				"error": "hobi not found",
			})
			return nil
		}

		if err != nil {
			panic(err)
		}

		_, err = mysql.Exec(`INSERT INTO mahasiswa_hobi
      VALUES ('','` + id + `' ,'` + hobi.ID + `' )`)
		if err != nil {
			panic(err)
		}

	}

	rows, err := mysql.Query("SELECT * FROM mahasiswa_hobi WHERE id_mahasiswa = " + id)
	var listHobi []string

	if err != nil {
		panic(err)
	}

	for rows.Next() {
		var mahasiswaHobi database.MahasiswaHobi

		err := rows.Scan(&mahasiswaHobi.ID, &mahasiswaHobi.IdMahasiswa, &mahasiswaHobi.IdHobi)
		if err != nil {
			panic(err)
		}

		row := mysql.QueryRow("SELECT * FROM hobi WHERE id = " + mahasiswaHobi.IdHobi)

		var hobi database.Hobi

		err = row.Scan(&hobi.ID, &hobi.NamaHobi)
		if hobi.ID != mahasiswaHobi.IdHobi {
			c.Status(404).JSON(&fiber.Map{
				"error": "hobi not found",
			})
			return nil
		}

		if err != nil {

			panic(err.Error())
		}

		listHobi = append(listHobi, hobi.NamaHobi)

	}

	mahasiswa.Hobi = listHobi
	mahasiswa.Jurusan = jurusan.NamaJurusan

	c.Status(200).JSON(&fiber.Map{
		"mahasiswa": mahasiswa,
	})
	return nil
}

func deleteMahasiswa(c *fiber.Ctx) error {

	id := c.Params("id")

	_, err := mysql.Exec(`DELETE FROM mahasiswa WHERE id = ` + id)
	if err != nil {
		c.Status(200).JSON(&fiber.Map{
			"error": err.Error(),
		})
	}

	_, err = mysql.Exec(`DELETE FROM mahasiswa_hobi WHERE id_mahasiswa = ` + id)
	if err != nil {
		c.Status(200).JSON(&fiber.Map{
			"error": err.Error(),
		})
	}

	c.Status(200).JSON(&fiber.Map{
		"message": "mahasiswa deleted successfully",
	})
	return nil
}

func SetupRoute(db *sql.DB) *fiber.App {
	app := *fiber.New()
	mysql = db
	app.Post("/api/v1/mahasiswa", createMahasiswa)
	app.Get("/api/v1/mahasiswa/:id", readMahasiswa)
	app.Get("/api/v1/mahasiswa", readMahasiswas)
	app.Put("/api/v1/mahasiswa/:id", updateMahasiswa)
	app.Delete("/api/v1/mahasiswa/:id", deleteMahasiswa)

	// jurusan
	app.Post("/api/v1/jurusan", createJurusan)
	app.Get("/api/v1/jurusan", readJurusans)

	// Hobi
	app.Post("/api/v1/hobi", createHobi)
	app.Get("/api/v1/hobi", readHobis)

	return &app
}
