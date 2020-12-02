package main

import (
	"database/sql"
	"fmt"
	"log"
	"math"

	_ "github.com/go-sql-driver/mysql"
)

type aa struct {
	id      int
	valor_a float64
}

type bb struct {
	id      int
	valor_b float64
}

type cc struct {
	id      int
	valor_c float64
}

type xxpos struct {
	id         int
	valor_xpos float64
}

type xxneg struct {
	id         int
	valor_xneg float64
}

func exec(db *sql.DB, sql string) sql.Result {
	result, err := db.Exec(sql)
	if err != nil {
		panic(err)
	}
	return result
}

func main() {

	db, err := sql.Open("mysql", "root:abcd1234@/")
	if err != nil {
		panic(err)
	}

	defer db.Close()

	var a float64
	var b float64
	var c float64
	var x1 float64
	var x2 float64
	codigo := 0
	codigo2 := 0
	codigo3 := 0
	codigo4 := 0
	codigo5 := 0
	codigo01 := 0
	codigo02 := 0
	codigo03 := 0
	codigo04 := 0
	codigo05 := 0
	var valoresa []float64
	var valoresb []float64
	var valoresc []float64
	var valoresxpos []float64
	var valoresxneg []float64
	totalregistros := 0
	totalregistros2 := 0
	totalregistros3 := 0
	totalregistros4 := 0
	totalregistros5 := 0
	valor01 := 0
	valor02 := 0
	valor03 := 0
	valor04 := 0
	valor05 := 0
	codigo10 := 0
	codigo20 := 0
	codigo30 := 0
	codigo40 := 0
	codigo50 := 0

	fmt.Println("\nCalculando uma equação de segundo grau")
	fmt.Printf("\n")

	fmt.Printf("Digite um valor A: ")
	fmt.Scanln(&a)

	fmt.Printf("Digite um valor B: ")
	fmt.Scanln(&b)

	fmt.Printf("Digite um valor C: ")
	fmt.Scanln(&c)

	// Calculando raizes
	x1 = (-b + math.Sqrt(b*b-4*a*c)) / (2 * a)

	x2 = (-b - math.Sqrt(b*b-4*a*c)) / (2 * a)

	fmt.Println("X+ = ", x1)
	fmt.Println("X- = ", x2)

	// Criando Database
	exec(db, "CREATE DATABASE IF NOT EXISTS tabelas")

	// Selecionando Database
	exec(db, "USE tabelas")

	// Criando tabela A
	exec(db, `CREATE TABLE IF NOT EXISTS a(
		valor_ida INT AUTO_INCREMENT PRIMARY KEY,
		valor_a DECIMAL(7,2)
	)`)

	// Criando tabela B
	exec(db, `CREATE TABLE IF NOT EXISTS b(
		valor_idb INT AUTO_INCREMENT PRIMARY KEY,
		valor_b DECIMAL(7,2)
	)`)

	// Criando tabela C
	exec(db, `CREATE TABLE IF NOT EXISTS c(
		valor_idc INT AUTO_INCREMENT PRIMARY KEY,
		valor_c DECIMAL(7,2)
	)`)

	// Criando tabela X+
	exec(db, `CREATE TABLE IF NOT EXISTS xpos(
		valor_idxpos INT AUTO_INCREMENT PRIMARY KEY,
		valor_xpos DECIMAL(7,2)
	)`)

	// Criando tabela X-
	exec(db, `CREATE TABLE IF NOT EXISTS xneg(
		valor_idxneg INT AUTO_INCREMENT PRIMARY KEY,
		valor_xneg DECIMAL(7,2)
	)`)

	// Adicionando valor A
	valoresa = append(valoresa, a)

	rows, _ := db.Query("SELECT MAX(valor_ida) AS ultimo FROM a")
	defer rows.Close()
	for rows.Next() {
		rows.Scan(&codigo)
		codigo = codigo + 1
	}
	totalregistros = len(valoresa)
	tx, _ := db.Begin()
	stmt, _ := tx.Prepare("INSERT INTO a(valor_ida, valor_a) VALUES(?,?)")

	for i := 0; i < totalregistros; i++ {
		stmt.Exec(codigo, valoresa[i])
		codigo = codigo + 1
	}

	if err != nil {
		tx.Rollback()
		log.Fatal(err)
	}
	tx.Commit()

	// Adicionando valor B
	valoresb = append(valoresb, b)

	rows, _ = db.Query("SELECT MAX(valor_idb) AS ultimo FROM b")
	defer rows.Close()
	for rows.Next() {
		rows.Scan(&codigo2)
		codigo2 = codigo2 + 1
	}
	totalregistros2 = len(valoresb)
	tx, _ = db.Begin()
	stmt, _ = tx.Prepare("INSERT INTO b(valor_idb, valor_b) values(?,?)")

	for i := 0; i < totalregistros2; i++ {
		stmt.Exec(codigo2, valoresb[i])
		codigo2 = codigo2 + 1
	}

	if err != nil {
		tx.Rollback()
		log.Fatal(err)
	}
	tx.Commit()

	// Adicionando valor C
	valoresc = append(valoresc, c)

	rows, _ = db.Query("SELECT MAX(valor_idc) AS ultimo FROM c")
	defer rows.Close()
	for rows.Next() {
		rows.Scan(&codigo3)
		codigo3 = codigo3 + 1
	}
	totalregistros3 = len(valoresc)
	tx, _ = db.Begin()
	stmt, _ = tx.Prepare("INSERT INTO c(valor_idc, valor_c) values(?,?)")

	for i := 0; i < totalregistros3; i++ {
		stmt.Exec(codigo3, valoresc[i])
		codigo3 = codigo3 + 1
	}

	if err != nil {
		tx.Rollback()
		log.Fatal(err)
	}
	tx.Commit()

	// Adicionando valor X+
	valoresxpos = append(valoresxpos, x1)

	rows, _ = db.Query("SELECT MAX(valor_idxpos) AS ultimo FROM xpos")
	defer rows.Close()
	for rows.Next() {
		rows.Scan(&codigo4)
		codigo4 = codigo4 + 1
	}
	totalregistros4 = len(valoresxpos)
	tx, _ = db.Begin()
	stmt, _ = tx.Prepare("INSERT INTO xpos(valor_idxpos, valor_xpos) VALUES(?,?)")

	for i := 0; i < totalregistros4; i++ {
		stmt.Exec(codigo4, valoresxpos[i])
		codigo4 = codigo4 + 1
	}

	if err != nil {
		tx.Rollback()
		log.Fatal(err)
	}
	tx.Commit()

	// Adicionando valor X-
	valoresxneg = append(valoresxneg, x2)

	rows, _ = db.Query("SELECT MAX(valor_idxneg) AS ultimo FROM xneg")
	defer rows.Close()
	for rows.Next() {
		rows.Scan(&codigo5)
		codigo5 = codigo5 + 1
	}
	totalregistros5 = len(valoresxneg)
	tx, _ = db.Begin()
	stmt, _ = tx.Prepare("INSERT INTO xneg(valor_idxneg, valor_xneg) VALUES(?,?)")

	for i := 0; i < totalregistros5; i++ {
		stmt.Exec(codigo5, valoresxneg[i])
		codigo5 = codigo5 + 1
	}

	if err != nil {
		tx.Rollback()
		log.Fatal(err)
	}
	tx.Commit()

	rows, _ = db.Query("SELECT valor_ida, valor_a FROM a")
	defer rows.Close()
	fmt.Println("\nAs seguintes variáveis A estao registradas:")
	for rows.Next() {
		var c aa
		rows.Scan(&c.id, &c.valor_a)
		fmt.Println(c)
	}

	rows, _ = db.Query("SELECT valor_idb, valor_b FROM b")
	defer rows.Close()
	fmt.Println("\nAs seguintes variáveis B estao registradas:")
	for rows.Next() {
		var c bb
		rows.Scan(&c.id, &c.valor_b)
		fmt.Println(c)
	}

	rows, _ = db.Query("SELECT valor_idc, valor_c FROM c")
	defer rows.Close()
	fmt.Println("\nAs seguintes variáveis C estao registradas:")
	for rows.Next() {
		var c cc
		rows.Scan(&c.id, &c.valor_c)
		fmt.Println(c)
	}

	rows, _ = db.Query("SELECT valor_idxpos, valor_xpos FROM xpos")
	defer rows.Close()
	fmt.Println("\nAs seguintes variáveis X+ estao registradas:")
	for rows.Next() {
		var c xxpos
		rows.Scan(&c.id, &c.valor_xpos)
		fmt.Println(c)
	}

	rows, _ = db.Query("SELECT valor_idxneg, valor_xneg FROM xneg")
	defer rows.Close()
	fmt.Println("\nAs seguintes variáveis X- estao registradas:")
	for rows.Next() {
		var c xxneg
		rows.Scan(&c.id, &c.valor_xneg)
		fmt.Println(c)
	}

	// Update A
	fmt.Print("Informe o código do valor A que deseja atualizar: ")
	fmt.Scanln(&codigo01)

	fmt.Print("Informe o novo valor : ")
	fmt.Scanln(&valor01)

	stmt, _ = db.Prepare("UPDATE a SET valor_a = ? WHERE valor_ida = ?")
	stmt.Exec(valor01, codigo01)

	// Update B
	fmt.Print("Informe o código do valor B que deseja atualizar: ")
	fmt.Scanln(&codigo02)

	fmt.Print("Informe o novo valor : ")
	fmt.Scanln(&valor02)

	stmt, _ = db.Prepare("UPDATE b SET valor_b = ? WHERE valor_idb = ?")
	stmt.Exec(valor02, codigo02)

	// Update C
	fmt.Print("Informe o código do valor C que deseja atualizar: ")
	fmt.Scanln(&codigo03)

	fmt.Print("Informe o novo valor : ")
	fmt.Scanln(&valor03)

	stmt, _ = db.Prepare("UPDATE c SET valor_c = ? WHERE valor_idc = ?")
	stmt.Exec(valor03, codigo03)

	// Update X+
	fmt.Print("Informe o código do valor X+ que deseja atualizar: ")
	fmt.Scanln(&codigo04)

	fmt.Print("Informe o novo valor : ")
	fmt.Scanln(&valor04)

	stmt, _ = db.Prepare("UPDATE xpos SET valor_xpos = ? WHERE valor_idxpos = ?")
	stmt.Exec(valor04, codigo04)

	// Update X-
	fmt.Print("Informe o código do valor X- que deseja atualizar: ")
	fmt.Scanln(&codigo05)

	fmt.Print("Informe o novo valor : ")
	fmt.Scanln(&valor05)

	stmt, _ = db.Prepare("UPDATE xneg SET valor_xneg = ? WHERE valor_idxneg = ?")
	stmt.Exec(valor05, codigo05)

	// Select A (Print)
	rows, _ = db.Query("SELECT valor_ida, valor_a FROM a")
	defer rows.Close()
	fmt.Println("\nAs seguintes variáveis A estao registradas:")
	for rows.Next() {
		var c aa
		rows.Scan(&c.id, &c.valor_a)
		fmt.Println(c)
	}

	// Select B (Print)
	rows, _ = db.Query("SELECT valor_idb, valor_b FROM b")
	defer rows.Close()
	fmt.Println("\nAs seguintes variáveis B estao registradas:")
	for rows.Next() {
		var c bb
		rows.Scan(&c.id, &c.valor_b)
		fmt.Println(c)
	}

	// Select C (Print)
	rows, _ = db.Query("SELECT valor_idc, valor_c FROM c")
	defer rows.Close()
	fmt.Println("\nAs seguintes variáveis C estao registradas:")
	for rows.Next() {
		var c cc
		rows.Scan(&c.id, &c.valor_c)
		fmt.Println(c)
	}

	// Select X+ (Print)
	rows, _ = db.Query("SELECT valor_idxpos, valor_xpos FROM xpos")
	defer rows.Close()
	fmt.Println("\nAs seguintes variáveis X+ estao registradas:")
	for rows.Next() {
		var c xxpos
		rows.Scan(&c.id, &c.valor_xpos)
		fmt.Println(c)
	}

	// Select X- (Print)
	rows, _ = db.Query("SELECT valor_idxneg, valor_xneg FROM xneg")
	defer rows.Close()
	fmt.Println("\nAs seguintes variáveis X- estao registradas:")
	for rows.Next() {
		var c xxneg
		rows.Scan(&c.id, &c.valor_xneg)
		fmt.Println(c)
	}

	// Delete A
	fmt.Println("Digite o código do valor A que deseja remover.")
	fmt.Scanln(&codigo10)

	stmt, _ = db.Prepare("DELETE FROM a WHERE valor_ida = ?")
	stmt.Exec(codigo10)

	// Delete B
	fmt.Println("Digite o código do valor B que deseja remover.")
	fmt.Scanln(&codigo20)

	stmt, _ = db.Prepare("DELETE FROM b WHERE valor_idb = ?")
	stmt.Exec(codigo20)

	// Delete C
	fmt.Println("Digite o código do valor C que deseja remover.")
	fmt.Scanln(&codigo30)

	stmt, _ = db.Prepare("DELETE FROM c WHERE valor_idc = ?")
	stmt.Exec(codigo30)

	// Delete X+
	fmt.Println("Digite o código do valor X+ que deseja remover.")
	fmt.Scanln(&codigo40)

	stmt, _ = db.Prepare("DELETE FROM xpos WHERE valor_idxpos = ?")
	stmt.Exec(codigo40)

	// Delete X-
	fmt.Println("Digite o código do valor X- que deseja remover.")
	fmt.Scanln(&codigo50)

	stmt, _ = db.Prepare("DELETE FROM xneg WHERE valor_idxneg = ?")
	stmt.Exec(codigo50)

	// Select A (Print)
	rows, _ = db.Query("SELECT valor_ida, valor_a FROM a")
	defer rows.Close()
	fmt.Println("\nAs seguintes variáveis A estao registradas:")
	for rows.Next() {
		var c aa
		rows.Scan(&c.id, &c.valor_a)
		fmt.Println(c)
	}

	// Select B (Print)
	rows, _ = db.Query("SELECT valor_idb, valor_b FROM b")
	defer rows.Close()
	fmt.Println("\nAs seguintes variáveis B estao registradas:")
	for rows.Next() {
		var c bb
		rows.Scan(&c.id, &c.valor_b)
		fmt.Println(c)
	}

	// Select C (Print)
	rows, _ = db.Query("SELECT valor_idc, valor_c FROM c")
	defer rows.Close()
	fmt.Println("\nAs seguintes variáveis C estao registradas:")
	for rows.Next() {
		var c cc
		rows.Scan(&c.id, &c.valor_c)
		fmt.Println(c)
	}

	// Select X+ (Print)
	rows, _ = db.Query("SELECT valor_idxpos, valor_xpos FROM xpos")
	defer rows.Close()
	fmt.Println("\nAs seguintes variáveis X+ estao registradas:")
	for rows.Next() {
		var c xxpos
		rows.Scan(&c.id, &c.valor_xpos)
		fmt.Println(c)
	}

	// Select X- (Print)
	rows, _ = db.Query("SELECT valor_idxneg, valor_xneg FROM xneg")
	defer rows.Close()
	fmt.Println("\nAs seguintes variáveis X- estao registradas:")
	for rows.Next() {
		var c xxneg
		rows.Scan(&c.id, &c.valor_xneg)
		fmt.Println(c)
	}
}
