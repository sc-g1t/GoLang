// Pacote do programa.
package main

// Bibliotecas
import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

// ???
func exec(db *sql.DB, sql string) sql.Result {
	result, err := db.Exec(sql)
	if err != nil {
		panic(err)
	}
	return result
}

func main() {
	// Abre e loga como root no mysql
	db, err := sql.Open("mysql", "root:abcd1234@/")
	if err != nil {
		panic(err)
	}

	// Fecha o Banco de Dados após a execução
	defer db.Close()

	// Cria Database webloja se já não existir
	fmt.Println("Criando banco de dados webloja.")
	exec(db, "CREATE DATABASE IF NOT EXISTS webloja")
	fmt.Println("Sucesso.")

	// Escolhe o Database webloja para ser usado
	fmt.Println("Selecionando banco de dados webloja.")
	exec(db, "USE webloja")
	fmt.Println("Sucesso.")

	// Criando tabelas
	// Cria tabela Lojas
	fmt.Println("Criando Tabela Lojas.")
	exec(db, `CREATE TABLE IF NOT EXISTS lojas(
		loja_id INT AUTO_INCREMENT PRIMARY KEY,
		loja VARCHAR(40),
		cidadeloja VARCHAR(40),
		estadoloja VARCHAR(40)
	)`)
	fmt.Println("Sucesso.")

	// Cria tabela Produtos
	fmt.Println("Criando Tabela Produtos.")
	exec(db, `CREATE TABLE IF NOT EXISTS produtos(
		produto_id INT AUTO_INCREMENT PRIMARY KEY,
		produto VARCHAR(40),
		quantidade INT NOT NULL,
		vencimento DATE NOT NULL,
		preco DECIMAL(20,2),
		custo DECIMAL(20,2)
	)`)
	fmt.Println("Sucesso.")

	// Cria tabela Estoque
	fmt.Println("Criando Tabela Estoque.")
	exec(db, `CREATE TABLE IF NOT EXISTS estoques(
		loja_id INT NOT NULL,
		produto_id INT NOT NULL,
		quantidade INT NOT NULL,
		PRIMARY KEY(loja_id, produto_id),
		FOREIGN KEY (loja_id) REFERENCES lojas(loja_id),
		FOREIGN KEY (produto_id) REFERENCES produtos(produto_id)
	)
`)
	fmt.Println("Sucesso.")

	// Cria tabela Categorias.
	fmt.Println("Criando tabela categorias.")
	exec(db, `CREATE TABLE IF NOT EXISTS categoria(
		categoria_id INT AUTO_INCREMENT PRIMARY KEY,
		categoria VARCHAR(40)
	)`)
	fmt.Println("Sucesso.")

	// Cria tabela Marcas.
	fmt.Println("Criando tabela marcas.")
	exec(db, `CREATE TABLE IF NOT EXISTS marcas(
		marcas_id INT AUTO_INCREMENT PRIMARY KEY,
		marca VARCHAR(40),
		paisorigem VARCHAR(40)
	)`)

	// Declaração de variáveis
	var categorias []string
	categoria := ""
	resposta := ""
	codigo := 0
	totalregistros := 0

	// Cria uma nova catégoria (Input do usuario)
	for {
		fmt.Println("Informe um nome para a categoria: ")
		fmt.Scanln(&categoria)
		categorias = append(categorias, categoria)

		fmt.Println("cadastrando a categoria ", categoria)

		fmt.Println("Deseja cadastrar uma nova categoria? s/n")
		fmt.Scanln(&resposta)

		if resposta != "s" {
			break
		}

		// ???
		rows, _ := db.Query("select max(categoria_id) as ultimo from categorias")
		defer rows.Close()
		for rows.Next() {
			rows.Scan(&codigo)
			codigo = codigo + 1
		}

		// ???
		totalregistros = len(categorias)

		// ???
		tx, _ := db.Begin()
		stmt, _ := tx.Prepare("insert into categorias(categoria_id, categoria) values(?,?)")

		// Faz o incremento no categoria_id ???
		for i := 0; i < totalregistros; i++ {
			stmt.Exec(codigo, categorias[i])
			codigo = codigo + 1
		}

		// ???
		if err != nil {
			tx.Rollback()
			log.Fatal(err)
		}
		tx.Commit()
	}

}
