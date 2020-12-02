package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

type produto struct {
	code  int
	nome  string
	preco float64
	qtde  int
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

	var nomes []string
	var precos []float64
	var qtdes []int64
	var nome string
	var preco float64
	var qtde int64
	codigo := 0
	totalregistros := 0
	resposta := ""

	// Criando Database
	exec(db, "CREATE DATABASE IF NOT EXISTS dbprodutos")
	fmt.Println("DB criado")

	// Selecionando Database
	exec(db, "USE dbprodutos")

	exec(db, `CREATE TABLE IF NOT EXISTS produtos(
			code INT AUTO_INCREMENT PRIMARY KEY,
			nome VARCHAR(20),
			preco DECIMAL(7,2),
			qtde INT
	)`)
	fmt.Println("Tabela produtos criada.")

	// ====		{Insert} 	====

	for {

		fmt.Println("Digite o nome do produto: ")
		fmt.Scanln(&nome)

		fmt.Println("Digite o preço do produto: ")
		fmt.Scanln(&preco)

		fmt.Println("Digite a quantidade do produto: ")
		fmt.Scanln(&qtde)

		nomes = append(nomes, nome)

		precos = append(precos, preco)

		qtdes = append(qtdes, qtde)

		fmt.Println("Deseja cadastrar outro produto? s/n")
		fmt.Scanln(&resposta)

		if resposta != "s" {
			break
		}

	}

	// Incremento do codigo/id
	rows, _ := db.Query("SELECT MAX(code) AS ultimo FROM produtos")
	defer rows.Close()
	for rows.Next() {
		rows.Scan(&codigo)
		codigo = codigo + 1
	}

	// Declara a ação de Inserir itens a Tabela
	totalregistros = len(nomes)
	tx, _ := db.Begin()
	stmt, _ := tx.Prepare("INSERT INTO produtos(code, nome, preco, qtde) values(?,?,?,?)")

	// Função que adiciona os produtos na tabela. [codigo, nomes, precos, qtdes]
	// ... substitui os VALUES(?,?,?,?)
	for i := 0; i < totalregistros; i++ {
		stmt.Exec(codigo, nomes[i], precos[i], qtdes[i])
		codigo = codigo + 1

	}

	// Erro
	if err != nil {
		tx.Rollback()
		log.Fatal(err)
	}
	// Commit, ou seja, efetua a ação
	tx.Commit()

	// ====		{Select}	====
	rows, _ = db.Query("SELECT code, nome, preco, qtde FROM produtos")
	defer rows.Close()
	fmt.Println("\nOs seguintes produtos estão registrados:")
	for rows.Next() {
		var c produto
		rows.Scan(&c.code, &c.nome, &c.preco, &c.qtde)
		fmt.Println(c)
	}

	// ==== 	{Update} 	====

	fmt.Print("Informe o código do Produto que deseja atualizar: ")
	fmt.Scanln(&codigo)

	fmt.Print("Informe o novo nome do Produto que deseja atualizar: ")
	fmt.Scanln(&nome)

	fmt.Println("Informe o novo preço do Produto: ")
	fmt.Scanln(&preco)

	fmt.Println("Informe a quantidade do Produto: ")
	fmt.Scanln(&qtde)

	// update
	stmt, _ = db.Prepare("UPDATE produtos SET nome = ?, preco = ?, qtde = ? WHERE code = ?")
	stmt.Exec(nome, preco, qtde, codigo)

	// ====		{Select}	====
	rows, _ = db.Query("SELECT code, nome, preco, qtde FROM produtos")
	defer rows.Close()
	fmt.Println("\nOs seguintes produtos estão registrados:")
	for rows.Next() {
		var c produto
		rows.Scan(&c.code, &c.nome, &c.preco, &c.qtde)
		fmt.Println(c)
	}

	// ====		{Delete}	====

	fmt.Println("Digite o código do Produto que deseja apagar.")
	fmt.Scanln(&codigo)

	stmt, _ = db.Prepare("DELETE FROM produtos WHERE code = ?")
	stmt.Exec(codigo)

	// ====		{Select}	====
	rows, _ = db.Query("SELECT code, nome, preco, qtde FROM produtos")
	defer rows.Close()
	fmt.Println("\nOs seguintes produtos estão registrados:")
	for rows.Next() {
		var c produto
		rows.Scan(&c.code, &c.nome, &c.preco, &c.qtde)
		fmt.Println(c)
	}

}
