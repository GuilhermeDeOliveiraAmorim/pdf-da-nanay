package main

import (
	"fmt"
	"log"

	"github.com/johnfercher/maroto/v2/pkg/components/text"
	"github.com/johnfercher/maroto/v2/pkg/consts/border"
	"github.com/johnfercher/maroto/v2/pkg/consts/fontstyle"
	"github.com/johnfercher/maroto/v2/pkg/consts/pagesize"

	"github.com/johnfercher/maroto/v2/pkg"
	"github.com/johnfercher/maroto/v2/pkg/config"
	"github.com/johnfercher/maroto/v2/pkg/props"
)

type Alpinista struct {
	Nome   string
	Cidade string
}

func main() {
	customFont := "arial-unicode-ms"
	customFontFile := "LondrinaSolid-Regular.ttf"

	repository := config.NewRepository().
		AddUTF8Font(customFont, fontstyle.Normal, customFontFile).
		AddUTF8Font(customFont, fontstyle.Italic, customFontFile).
		AddUTF8Font(customFont, fontstyle.Bold, customFontFile).
		AddUTF8Font(customFont, fontstyle.BoldItalic, customFontFile)

	builder, err := config.NewBuilder().TryLoadRepository(repository)
	if err != nil {
		log.Fatal(err.Error())
	}

	cfg := builder.WithDefaultFont(&props.Font{Family: customFont}).
		WithPageSize(pagesize.Nanay).
		WithMargins(5, 13, 5).
		WithDebug(true).
		WithDimensions(215.9, 287.1).
		Build()

	alpinistas := []Alpinista{}

	for _, d := range getContents() {
		pessoa := Alpinista{
			Nome:   d[0],
			Cidade: d[1],
		}
		alpinistas = append(alpinistas, pessoa)
	}

	for _, d := range alpinistas {
		fmt.Println(d.Nome)
	}

	mrt := pkg.NewMaroto(cfg)
	m := pkg.NewMetricsDecorator(mrt)

	colStyle := &props.Cell{
		BorderType:  border.Full,
		BorderColor: &props.Color{Red: 255, Green: 255, Blue: 255},
	}

	rowStyle := &props.Cell{
		BorderColor: &props.Color{Red: 255, Green: 255, Blue: 255},
	}

	for _, alpinista := range alpinistas {
		m.AddRow(12.7,
			text.NewCol(4, alpinista.Nome, props.Text{Size: 18, Top: 5}).WithStyle(colStyle),
			text.NewCol(4, alpinista.Nome, props.Text{Size: 18, Top: 5}).WithStyle(colStyle),
			text.NewCol(4, alpinista.Nome, props.Text{Size: 18, Top: 5}).WithStyle(colStyle),
		).WithStyle(rowStyle)

		m.AddRow(12.7,
			text.NewCol(4, alpinista.Cidade, props.Text{Size: 14}).WithStyle(colStyle),
			text.NewCol(4, alpinista.Cidade, props.Text{Size: 14}).WithStyle(colStyle),
			text.NewCol(4, alpinista.Cidade, props.Text{Size: 14}).WithStyle(colStyle),
		).WithStyle(rowStyle)
	}

	document, err := m.Generate()
	if err != nil {
		log.Fatal(err.Error())
	}

	err = document.Save("etiquetas-anual-aju-2023.pdf")
	if err != nil {
		log.Fatal(err.Error())
	}
}

func getContents() [][]string {
	return [][]string{
		{"Deralds", "Canavieiras"},
		{"Edson", "Canavieiras"},
		{"JG", "Canavieiras"},
		{"Têka", "Canavieiras"},
		{"Tekas", "Canavieiras"},
		{"Tephy", "Canavieiras"},
		{"Damacenu", "Canavieiras"},
		{"Piu", "Canavieiras"},
		{"Dan Silva", "Canavieiras"},
		{"Naninha", "Juazeiro"},
		{"Ruh", "Juazeiro"},
		{"Drico", "Salvador"},
		{"Cascão", "Salvador"},
		{"Bia", "Salvador"},
		{"Elma", "Salvador"},
		{"Kinha", "Salvador"},
		{"Tobi", "Salvador"},
		{"Gabs", "Salvador"},
		{"JoMa", "Salvador"},
		{"Zé", "Salvador"},
		{"Lavi", "Salvador"},
		{"Lucas", "Salvador"},
		{"Harry", "Salvador"},
		{"Rêe", "Salvador"},
		{"Teté", "Salvador"},
		{"Thaci", "Salvador"},
		{"Milla", "Salvador"},
		{"Thiaguinho", "Salvador"},
		{"Bibia", "Salvador"},
		{"Majô", "Salvador"},
		{"Pams", "Salvador"},
		{"Jojô", "Salvador"},
		{"Maury Vitor", "Canavieiras"},
		{"Gustavao", "Salvador"},
		{"Gugu", "Canavieiras"},
		{"Lu Lyrio", "Cruz das Almas"},
		{"Lyriozinha", "Cruz das Almas"},
		{"Alana", "Aracaju"},
		{"Aninha", "Aracaju"},
		{"Lelê", "Aracaju"},
		{"Pablo", "Aracaju"},
		{"Vih", "Aracaju"},
		{"Lari", "Aracaju"},
		{"João Reis", "Juazeiro"},
		{"João", "Aracaju"},
		{"Kel", "Aracaju"},
		{"Ju Gregório", "Aracaju"},
		{"Leozinho", "Aracaju"},
		{"Flavinha", "Aracaju"},
		{"Mariana Dias", "Cruz das Almas"},
		{"Lu", "Aracaju"},
		{"Gabi Menezes", "Aracaju"},
		{"Nath", "Aracaju"},
		{"Nanna", "Aracaju"},
		{"Nanda Borges", "Juazeiro"},
		{"Livinha", "Aracaju"},
		{"Nilzinha", "Juazeiro"},
		{"Tio Nilton", "Aracaju"},
		{"Elô", "Aracaju"},
		{"Same", "Aracaju"},
		{"Tia Gilma", "Aracaju"},
		{"Eliza", "Juazeiro"},
		{"Esquerdinha", "Canavieiras"},
		{"Bonito", "Canavieiras"},
		{"Rafa Sobral", "Salvador"},
		{"Lua", "Salvador"},
		{"Thi", "Salvador"},
		{"Bebela", "Aracaju"},
		{"Ká", "Juazeiro"},
		{"Luquinhas", "Aracaju"},
		{"Duda Fonseca", "Cruz das Almas"},
		{"Paulinha", "Juazeiro"},
		{"Livinha", "Canavieiras"},
		{"FÁ", "Una"},
		{"Bia Rocha", "Canavieiras"},
		{"Tia Rita", "Canavieiras"},
		{"Sambras", "Canavieiras"},
		{"Juju", "Canavieiras"},
		{"Duda Fernandes", "Salvador"},
		{"Tchengo", "Cruz das Almas"},
		{"Mari Guimarães", "Aracaju"},
		{"Annie", "Aracaju"},
		{"Ju Secundo", "Aracaju"},
		{"Fabi", "Salvador"},
		{"Duda Alves", "Juazeiro"},
		{"Paulo", "Aracaju"},
		{"Neto", "Juazeiro"},
		{"Mairinha", "Aracaju"},
		{"Tati", "Aracaju"},
		{"Vivi", "Salvador"},
		{"Tuco", "Salvador"},
		{"Jean", "Itabuna"},
		{"Fran", "Itabuna"},
		{"Gabi", "Aracaju"},
		{"Juninho", "Aracaju"},
		{"Nay", "Aracaju"},
		{"Kiko", "Aracaju"},
		{"Gabi", "Salvador"},
		{"Ray", "Salvador"},
		{"Nat", "Salvador"},
		{"Cathy", "Salvador"},
		{"Bê", "Itabuna"},
		{"Leo Barros", "Aracaju"},
		{"Bia", "Aracaju"},
		{"Lis", "Jequié"},
		{"Fuska", "Aracaju"},
		{"Diana", "Aracaju"},
		{"Nath Dantas", "Salvador"},
		{"Smurfette", "Jequié"},
		{"Arthur", "Jequié"},
		{"Bruno", "Jequié"},
		{"MMMM", "Jequié"},
		{"Carlinha", "Salvador"},
		{"Letícia", "Salvador"},
		{"Rod", "Salvador"},
		{"Lua", "Jequié"},
		{"Tali", "Aracaju"},
		{"Davi", "Aracaju"},
		{"Rafael", "Jequié"},
		{"Ana", "Jequié"},
		{"Lu", "Itabuna"},
		{"Carina", "Aracaju"},
		{"Pedroca", "Aracaju"},
		{"Brunão", "Aracaju"},
		{"Vitorinha", "Aracaju"},
		{"Carol", "Aracaju"},
		{"Alex", "Aracaju"},
		{"Lana", "Aracaju"},
		{"Mario", "Jequié"},
		{"Bih Aguiar", "Aracaju"},
		{"San", "Jequié"},
		{"Luarinha", "Jequié"},
		{"Canário", "Jequié"},
		{"Tia Alê", "Aracaju"},
		{"Bella", "Aracaju"},
		{"Yas", "Jequié"},
		{"Ratinho", "Jequié"},
		{"Lu", "Aracaju"},
		{"Sams", "Aracaju"},
		{"Larinha", "Aracaju"},
		{"Lê", "Una"},
		{"Ana Lu", "Una"},
		{"Gui", "Aracaju"},
		{"Su", "Cruz das Almas"},
		{"Hillis", "Jequié"},
		{"Bruninho", "Jequié"},
		{"Ellen", "Jequié"},
		{"Xande", "Jequié"},
		{"Irisinha", "Aracaju"},
		{"Boca", "Salvador"},
		{"Lica", "Jequié"},
		{"Lorrany", "Jequié"},
		{"Fadinha", "Salvador"},
		{"Lalas", "Aracaju"},
		{"Driquinha", "Aracaju"},
		{"Sté", "Jequié"},
		{"Ita", "Salvador"},
		{"Mari Moscoso", "Aracaju"},
		{"Ju", "Aracaju"},
		{"Lari Santos", "Aracaju"},
		{"May Coelho", "Salvador"},
		{"Lelê", "Jequié"},
		{"Helder", "Aracaju"},
		{"Ju", "Aracaju"},
		{"Thi", "Jequié"},
		{"Letícia", "Jequié"},
		{"Gabi", "Jequié"},
		{"Matheus", "Aracaju"},
		{"Iury", "Jequié"},
		{"Beca", "Jequié"},
		{"Bibia", "Jequié"},
		{"Juli", "Aracaju"},
		{"Sophi", "Salvador"},
		{"Manu Vilanova", "Aracaju"},
		{"Gui", "Jequié"},
		{"Lua", "Jequié"},
		{"Peu", "Salvador"},
		{"Mica", "Jequié"},
		{"Vic", "Canavieiras"},
		{"Tio Brinquinho", "Aracaju"},
		{"João", "Aracaju"},
		{"Cris", "Cruz das Almas"},
		{"Novais", "Aracaju"},
		{"Léo Hora", "Aracaju"},
		{"Pedroca", "Canavieiras"},
		{"Mateuzinho", "Aracaju"},
		{"Rick", "Jequié"},
		{"Anna", "Salvador"},
		{"Tw", "Canavieiras"},
		{"Nat Alumiar", "Salvador"},
		{"Lau", "Cruz das Almas"},
		{"Nina", "Jequié"},
		{"Bel", "Aracaju"},
		{"Antônio", "Salvador"},
		{"Nanda Freitas", "Salvador"},
		{"Bela", "Cruz das Almas"},
		{"Neto", "Aracaju"},
		{"Dirce", "Aracaju"},
		{"Karlinha P.", "Aracaju"},
		{"Grazi Moreno", "Salvador"},
		{"Lara Seabra", "Aracaju"},
		{"Super Choque", "Jequié"},
		{"Biscoito", "Aracaju"},
		{"Ju Gentil", "Aracaju"},
		{"Lore Maria", "Aracaju"},
		{"My", "Aracaju"},
		{"Nena", "Aracaju"},
		{"Laís", "Aracaju"},
		{"Drico", "Juazeiro"},
		{"Mile", "Salvador"},
		{"Carlinha", "Juazeiro"},
		{"Maria", "Aracaju"},
		{"Raissa", "Aracaju"},
		{"Miguel", "Jequié"},
		{"Tutu", "Jequié"},
		{"Tia Soeli", "Aracaju"},
		{"Gabi", "Aracaju"},
		{"Tio Mauro", "Aracaju"},
		{"Kati", "Aracaju"},
		{"Max", "Aracaju"},
		{"Lelê", "Petrolina"},
		{"Vic", "Petrolina"},
		{"Pedro Sampaio", "Aracaju"},
		{"Marcelo", "Aracaju"},
		{"Neto", "Petrolina"},
		{"Agroboy", "Petrolina"},
		{"Tia Lu", "Aracaju"},
		{"Rafha", "Petrolina"},
		{"Jana", "Aracaju"},
		{"Tio Gilton", "Aracaju"},
		{"Suanne", "Petrolina"},
		{"Paulyran", "Petrolina"},
		{"Vi Marcula", "Petrolina"},
		{"Garibaldi", "Petrolina"},
		{"Serra", "Petrolina"},
		{"Lucas", "Petrolina"},
		{"Milla Faria", "Petrolina"},
		{"Uel", "Petrolina"},
		{"Fonseca", "Salvador"},
		{"Franci", "Una"},
		{"Thays", "Petrolina"},
		{"Mateus", "Petrolina"},
		{"Ceará", "Petrolina"},
		{"Kamille", "Petrolina"},
		{"Babi", "Petrolina"},
		{"Carol", "Petrolina"},
		{"Lore", "Petrolina"},
		{"Celestino", "Petrolina"},
		{"Gabi", "Petrolina"},
		{"Lelê Gurgel", "Petrolina"},
		{"Tia Miriam", "Petrolina"},
		{"Mila", "Petrolina"},
		{"Van", "Aracaju"},
		{"Cadu Maverick", "Petrolina"},
		{"Tia Neuri", "Petrolina"},
		{"Alice", "Petrolina"},
		{"Biel Maciel", "Petrolina"},
		{"Bina", "Petrolina"},
		{"Jupira", "Salvador"},
		{"Cassinha", "Petrolina"},
		{"Pedrinho", "Aracaju"},
		{"Mavi Gomes", "Petrolina"},
		{"Thata", "Petrolina"},
		{"Ferraz", "Petrolina"},
		{"Biel Castro", "Petrolina"},
		{"Lari", "Salvador"},
		{"Tina", "Aracaju"},
		{"Alice", "Petrolina"},
		{"Lu Mitsuyo", "Petrolina"},
		{"Emanuel", "Petrolina"},
		{"Wolverine", "Petrolina"},
		{"Carol Serafim", "Petrolina"},
		{"Zezero", "Petrolina"},
		{"Felipe Rôni", "Petrolina"},
		{"Gui", "Aracaju"},
		{"Lua", "Petrolina"},
		{"Mateus", "Petrolina"},
		{"Ray", "Petrolina"},
		{"Lukinhas", "Petrolina"},
		{"Guind", "Petrolina"},
		{"Verenna", "Petrolina"},
		{"Biguinho", "Petrolina"},
		{"Angela", "Petrolina"},
		{"Pequeno", "Petrolina"},
		{"Ju", "Ilhéus"},
		{"Tia Fátima", "Petrolina"},
		{"Luiz", "Ilhéus"},
		{"Tia Berna", "Petrolina"},
		{"Ryan", "Ilhéus"},
		{"Elinha", "Aracaju"},
		{"Higor", "Ilhéus"},
		{"Juce", "Salvador"},
		{"Mano Gusta", "Ilhéus"},
		{"Lua", "Ilhéus"},
		{"Anderson", "Petrolina"},
		{"Tacy", "Ilhéus"},
		{"Jhonnys", "Petrolina"},
		{"Biel Ribeiro", "Petrolina"},
		{"Duda", "Ilhéus"},
		{"Tio Renan", "Petrolina"},
		{"Gabi Matias", "Petrolina"},
		{"Paulinho", "Aracaju"},
		{"Maick", "Petrolina"},
		{"Alida", "Petrolina"},
		{"Dani", "Aracaju"},
		{"Nanda", "Aracaju"},
		{"Lice", "Petrolina"},
		{"Mi", "Petrolina"},
		{"Adson", "Petrolina"},
		{"Lua", "Canavieiras"},
		{"Corujinha", "Petrolina"},
		{"Thai", "Cruz das Almas"},
		{"Beca", "Petrolina"},
		{"Paula Nogueira", "Petrolina"},
		{"Gilson", "Aracaju"},
		{"Josiely", "Petrolina"},
		{"Guy", "Aracaju"},
		{"JL", "Aracaju"},
		{"Bebê", "Canavieiras"},
		{"Helô", "Cruz das Almas"},
		{"Bia Pacheco", "Petrolina"},
		{"Mari Accioly", "Aracaju"},
		{"Pacheco", "Petrolina"},
		{"Bia", "Lauro de Freitas"},
		{"Babi", "Lauro de Freitas"},
		{"Mafê", "Lauro de Freitas"},
		{"Duda", "Salvador"},
		{"Livia", "Petrolina"},
		{"Ju", "Salvador"},
		{"Pimenta", "Lauro de Freitas"},
		{"Yudi", "Petrolina"},
		{"Lucas", "Lauro de Freitas"},
		{"Senhor", "Lauro de Freitas"},
		{"Synara", "Petrolina"},
		{"Lari Cardoso", "Aracaju"},
		{"Sté", "Salvador"},
		{"Amanda Soares", "Petrolina"},
		{"Jacque", "Lauro de Freitas"},
		{"Luan", "Petrolina"},
		{"Gabi", "Lauro de Freitas"},
		{"Nana", "Aracaju"},
		{"Vivi", "Lauro de Freitas"},
		{"Tali", "Salvador"},
		{"Vic", "Lauro de Freitas"},
		{"Tio Manoel", "Petrolina"},
		{"Let", "Lauro de Freitas"},
		{"Bru", "Lauro de Freitas"},
		{"Raica", "Aracaju"},
		{"Alê", "Aracaju"},
		{"Sam Martins", "Petrolina"},
		{"Coroinha", "Petrolina"},
		{"Hellen", "Lauro de Freitas"},
		{"Nando", "Lauro de Freitas"},
		{"Carla", "Petrolina"},
		{"Julia Maria", "Petrolina"},
		{"Lari", "Salvador"},
		{"Fabi", "Lauro de Freitas"},
		{"Let Cardoso", "Aracaju"},
		{"Gui", "Lauro de Freitas"},
		{"Kinhos", "Lauro de Freitas"},
		{"Helo", "Petrolina"},
		{"Pionorio", "Petrolina"},
		{"Paulo", "Lauro de Freitas"},
		{"Doca", "Petrolina"},
		{"Lucca", "Aracaju"},
		{"Memy", "Lauro de Freitas"},
		{"Amanda", "Lauro de Freitas"},
		{"Sol", "Lauro de Freitas"},
		{"Juju Gomes", "Petrolina"},
		{"Camis", "Lauro de Freitas"},
		{"Paixão", "Lauro de Freitas"},
		{"Pri", "Lauro de Freitas"},
		{"Arthur 99", "Petrolina"},
		{"Juh Leite", "Petrolina"},
		{"Cela", "Lauro de Freitas"},
		{"Regi", "Lauro de Freitas"},
		{"Peu", "Lauro de Freitas"},
		{"Duda Mendes", "Petrolina"},
		{"Fróes", "Lauro de Freitas"},
		{"Juju", "Lauro de Freitas"},
		{"Johannes", "Petrolina"},
		{"Mari Muniz", "Lauro de Freitas"},
		{"Biel simoes", "Petrolina"},
		{"Dudinha", "Lauro de Freitas"},
		{"Samarinha", "Petrolina"},
		{"Cauã", "Lauro de Freitas"},
		{"Marinho", "Lauro de Freitas"},
		{"Nati", "Lauro de Freitas"},
		{"Malu", "Petrolina"},
		{"Oswaldo", "Lauro de Freitas"},
		{"Chan", "Petrolina"},
		{"Edu", "Lauro de Freitas"},
		{"João Lima", "Lauro de Freitas"},
		{"Lala", "Petrolina"},
		{"Lou", "Lauro de Freitas"},
		{"Bazante", "Petrolina"},
		{"Jujuba", "Lauro de Freitas"},
		{"Murilo Barbosa", "Petrolina"},
		{"Dora", "Petrolina"},
		{"Claudia", "Petrolina"},
		{"Kiki", "Petrolina"},
		{"Mamá Rodrigues", "Petrolina"},
		{"Ionara", "Petrolina"},
		{"Tatá Ribeiro", "Petrolina"},
		{"Felipinho", "Petrolina"},
		{"Alissuu", "Petrolina"},
		{"Bea", "Petrolina"},
		{"Jojo", "Petrolina"},
		{"Sanchez", "Petrolina"},
		{"Jotinha", "Petrolina"},
		{"Gui", "Petrolina"},
		{"Murilo Coelho", "Petrolina"},
		{"Carlinha", "Cruz das Almas"},
		{"Netinho", "Cruz das Almas"},
		{"Juh Brito", "Petrolina"},
		{"Clecielly", "Petrolina"},
		{"Lyla", "Petrolina"},
		{"Bruno", "Petrolina"},
		{"Marcinha", "Petrolina"},
		{"Sabrina", "Petrolina"},
		{"Catita", "Petrolina"},
		{"Lay", "Petrolina"},
		{"Tonhão", "Petrolina"},
		{"Paulinha", "Petrolina"},
		{"Tia Margarida", "Aracaju"},
		{"Zagueiro", "Petrolina"},
		{"Pedro", "Petrolina"},
		{"Flaris", "Petrolina"},
		{"Rosinha", "Aracaju"},
		{"Mandy", "Petrolina"},
		{"Kiki", "Petrolina"},
		{"Val", "Petrolina"},
		{"Nina", "Aracaju"},
		{"Amauri", "Petrolina"},
		{"Serginho", "Aracaju"},
		{"Glenison", "Petrolina"},
		{"Vaqueiro ", "Petrolina"},
		{"Thiaguinho", "Petrolina"},
		{"Dudu", "Aracaju"},
		{"Jessé", "Petrolina"},
		{"Laura", "Petrolina"},
		{"Marcelo", "Salvador"},
		{"Iego Viana", "Petrolina"},
		{"Pallominha", "Juazeiro"},
		{"Lila", "Petrolina"},
		{"Mari Gomes", "Petrolina"},
		{"Aninha", "Petrolina"},
		{"Carlinha", "Petrolina"},
		{"Daniel", "Petrolina"},
		{"Meiry", "Petrolina"},
		{"Bia Carvalho", "Petrolina"},
		{"Mimi", "Petrolina"},
		{"Fernando", "Petrolina"},
		{"Vini Guirra", "Petrolina"},
		{"Conras", "Petrolina"},
		{"Totó", "Petrolina"},
		{"Fonseca", "Petrolina"},
		{"Jossandra", "Petrolina"},
		{"Avelino", "Petrolina"},
		{"Thau", "Petrolina"},
		{"Kassinha", "Petrolina"},
		{"Ray Pilé", "Petrolina"},
		{"Bru", "Petrolina"},
		{"Gio", "Petrolina"},
		{"Almeida", "Petrolina"},
		{"Clarinha", "Petrolina"},
		{"Tia Débora", "Petrolina"},
		{"Vitinho", "Salvador"},
		{"Lari Carvalho", "Petrolina"},
		{"Leninha", "Petrolina"},
		{"Zigu", "Petrolina"},
		{"Isa Lopes", "Petrolina"},
		{"B.A", "Canavieiras"},
		{"Lore", "Salvador"},
		{"Paula", "Salvador"},
		{"Dany", "Cruz das Almas"},
		{"Lila", "Aracaju"},
		{"Jimmy", "Aracaju"},
		{"Lore Couto", "Aracaju"},
		{"Yargo", "Aracaju"},
		{"Cecília Pithon", "Jequié"},
		{"Culas", "Itabuna"},
		{"Nati", "Jacobina"},
		{"Cunha", "Jacobina"},
		{"Mila", "Jacobina"},
		{"Mi", "Jacobina"},
		{"Laranjeira", "Jacobina"},
		{"Blu", "Jacobina"},
		{"Gonzaga", "Itabuna"},
		{"Fifi", "Jacobina"},
		{"Rafa Souza", "Jacobina"},
		{"Isa", "Jacobina"},
		{"Naecinho", "Jacobina"},
		{"Isis", "Jacobina"},
		{"Choco", "Jacobina"},
		{"Mari Amorim", "Jacobina"},
		{"Ana Bia", "Jacobina"},
		{"Clara", "Jacobina"},
		{"Gaby Goes", "Jacobina"},
		{"Gio", "Jacobina"},
		{"JP", "Jacobina"},
		{"Dono do Escalada", "Jacobina"},
		{"Peu", "Lauro de Freitas"},
		{"Lipe", "Petrolina"},
		{"Mel", "Petrolina"},
		{"Wel", "Aracaju"},
		{"Bobzera", "Canavieiras"},
		{"Dez", "Jacobina"},
		{"Lulu", "Jacobina"},
		{"Bela", "Jacobina"},
		{"Neres", "Jacobina"},
		{"Isa", "Jacobina"},
		{"Didi", "Jacobina"},
		{"Pe. Seni", "Jacobina"},
		{"Procópio", "Jacobina"},
		{"Tia Nadja", "Jacobina"},
		{"Mari", "Jacobina"},
		{"Vânia", "Jacobina"},
		{"Mima", "Jacobina"},
		{"Lore", "Jacobina"},
		{"Alicinha", "Jacobina"},
		{"Lipe", "Jacobina"},
		{"Yah", "Aracaju"},
		{"Iza", "Juazeiro"},
		{"Maria Gabi", "Jacobina"},
		{"Tia Lila", "Jacobina"},
		{"Teteu", "Jacobina"},
		{"Lay", "Jacobina"},
		{"Ju Moreira", "Jacobina"},
		{"Silva", "Jacobina"},
		{"Valquinho", "Jacobina"},
		{"Kiki", "Jacobina"},
		{"Alemão", "Jacobina"},
		{"Hugo", "Jacobina"},
		{"Duda F", "Jacobina"},
		{"Nicinho", "Jacobina"},
		{"Ju Ribeiro", "Jacobina"},
		{"Duda Dias", "Jacobina"},
		{"Rafa Leah", "Jacobina"},
		{"Gogs", "Jacobina"},
		{"Ursinho", "Jacobina"},
		{"Tia Déa", "Jacobina"},
		{"Atani", "Jacobina"},
		{"Jhon", "Jequié"},
		{"Filipe", "Aracaju"},
		{"Vi", "Itabuna"},
		{"Lariii", "Itabuna"},
		{"Davi", "Jequié"},
		{"Pedro", "Aracaju"},
		{"Lalinha", "Juazeiro"},
		{"Caxin", "Jequié"},
		{"Lavi", "Jacobina"},
		{"Nique", "Jacobina"},
		{"Levy", "Jacobina"},
		{"Belly", "Jacobina"},
		{"Black", "Jacobina"},
		{"Cainho", "Jacobina"},
		{"Tia Slene", "Juazeiro"},
		{"Rai", "Jacobina"},
		{"Biel", "Jacobina"},
		{"Tia Cléa", "Jacobina"},
		{"Ste", "Itabuna"},
		{"Kaká", "Jacobina"},
		{"Bê", "Jacobina"},
		{"Belém", "Canavieiras"},
		{"Hazi", "Jacobina"},
		{"Mah", "Jacobina"},
		{"Colágeno", "Jacobina"},
		{"Paiva", "Jacobina"},
		{"Lou", "Jacobina"},
		{"Quel", "Jacobina"},
		{"Neto", "Jacobina"},
		{"Bino", "Cruz das Almas"},
		{"Karlinha Dani", "Aracaju"},
		{"Nick", "Jequié"},
		{"Bruno", "Jequié"},
		{"Cananda", "Cruz das Almas"},
		{"Ludi", "Itabuna"},
		{"Nanda", "Itabuna"},
		{"Ray", "Cruz das Almas"},
		{"Ru", "Jequié"},
		{"Mine", "Itabuna"},
		{"Winstinho", "Aracaju"},
		{"Bia", "Cruz das Almas"},
		{"Clariana", "Cruz das Almas"},
		{"Leti", "Aracaju"},
		{"Tia Regi", "Jacobina"},
		{"Claudinho", "Cruz das Almas"},
		{"Lelê", "Aracaju"},
		{"Xico", "Lauro de Freitas"},
		{"Tia Bia", "Lauro de Freitas"},
		{"Sofi", "Lauro de Freitas"},
		{"Jô", "Aracaju"},
		{"Isa de Maria", "Itabuna"},
		{"Tia Luana", "Aracaju"},
		{"Tio Léo", "Aracaju"},
		{"Marcelo", "Aracaju"},
		{"Robertinho", "Aracaju"},
		{"Melancia", "Canavieiras"},
		{"Nique", "Itabuna"},
		{"Yagaum", "Itabuna"},
		{"Mari Cavalcante", "Aracaju"},
		{"Tia Neide", "Juazeiro"},
		{"Alex", "Juazeiro"},
		{"Lai", "Aracaju"},
		{"Pepeu", "Juazeiro"},
		{"Kibe", "Itabuna"},
		{"A Loka", "Juazeiro"},
		{"Jana", "Juazeiro"},
		{"Lucigol", "Jequié"},
		{"Isa Melo", "Aracaju"},
		{"lalá", "Canavieiras"},
		{"El Potro", "Juazeiro"},
		{"Tililinho", "Cruz das Almas"},
		{"Lelê", "Canavieiras"},
		{"Thatai", "Canavieiras"},
		{"Rasta", "Canavieiras"},
		{"Marcelo", "Aracaju"},
		{"Julia C", "Aracaju"},
		{"Rafa Costa", "Canavieiras"},
		{"Di Caprio", "Aracaju"},
		{"Clarinha", "Aracaju"},
		{"Liana", "Aracaju"},
		{"Milinha", "Aracaju"},
		{"Tiago", "Aracaju"},
		{"Paulisson", "Aracaju"},
		{"Maly", "Aracaju"},
		{"Nah", "Itabuna"},
		{"Tio Chico", "Aracaju"},
		{"Tia Noelia", "Aracaju"},
		{"Ianna", "Canavieiras"},
		{"Karol", "Canavieiras"},
		{"Mile Lessa", "Cruz das Almas"},
		{"Caioh", "Conceição do Almeida"},
		{"Victor Oasis", "Aracaju"},
		{"Sther", "Conceição do Almeida"},
		{"Meizã", "Conceição do Almeida"},
		{"Jú", "Itabuna"},
		{"Yago Alcantara", "Conceição do Almeida"},
		{"Thi", "Conceição do Almeida"},
		{"Dai", "Conceição do Almeida"},
		{"Chris", "Aracaju"},
		{"Biih", "Conceição do Almeida"},
		{"Tia Val", "Conceição do Almeida"},
		{"Duda Lima", "Conceição do Almeida"},
		{"Tia Sônia", "Conceição do Almeida"},
		{"Whaguinho", "Conceição do Almeida"},
		{"KT", "Aracaju"},
		{"Tia Déa", "Conceição do Almeida"},
		{"Marisol", "Canavieiras"},
		{"Lais", "Aracaju"},
		{"Tia Conce", "Conceição do Almeida"},
		{"Tio Zeca", "Aracaju"},
		{"Tia Anita", "Aracaju"},
		{"Tia Mariza", "Conceição do Almeida"},
		{"Tio Luciano", "Conceição do Almeida"},
		{"Cat", "Itabuna"},
		{"Nando", "Cruz das Almas"},
		{"Tia Maria", "Conceição do Almeida"},
		{"Lipe", "Cruz das Almas"},
		{"Jope", "Aracaju"},
		{"Tia Kátia", "Aracaju"},
		{"Ju", "Cruz das Almas"},
		{"Bia", "Aracaju"},
		{"Karlinha", "Itabuna"},
		{"Tia Sandra", "Aracaju"},
		{"Bely", "Aracaju"},
		{"Babalu", "Itabuna"},
		{"Fernando", "Aracaju"},
		{"Tia Maruse", "Itabuna"},
		{"Índia", "Cruz das Almas"},
		{"Duda", "Itabuna"},
		{"Zu", "Aracaju"},
		{"Thi", "Una"},
		{"Flores", "Cruz das Almas"},
		{"Naná", "Aracaju"},
		{"Dudu", "Itabuna"},
		{"Clarinha", "Cruz das Almas"},
		{"Tio Hamilton", "Conceição do Almeida"},
		{"Vitor Lessa", "Cruz das Almas"},
		{"Tia Cida", "Conceição do Almeida"},
		{"Miva", "Cruz das Almas"},
		{"Manu", "Aracaju"},
		{"Tia Josy", "Conceição do Almeida"},
		{"Mi Souza", "Cruz das Almas"},
		{"Gadé", "Petrolina"},
		{"Tia Carioca", "Aracaju"},
		{"Guto", "Conceição do Almeida"},
		{"Dessa", "Conceição do Almeida"},
		{"Kaká", "Conceição do Almeida"},
		{"Diego", "Conceição do Almeida"},
		{"Zay", "Conceição do Almeida"},
		{"Gabyy", "Conceição do Almeida"},
		{"Ray", "Conceição do Almeida"},
		{"Gui", "Conceição do Almeida"},
		{"Lalai", "Conceição do Almeida"},
		{"Ga", "Conceição do Almeida"},
		{"Éto", "Canavieiras"},
		{"Maria Gio", "Juazeiro"},
		{"Lulu", "Juazeiro"},
		{"Bia", "Conceição do Almeida"},
		{"Tchui", "Conceição do Almeida"},
		{"Teca", "Conceição do Almeida"},
		{"Dourado", "Conceição do Almeida"},
		{"Bia", "Conceição do Almeida"},
		{"Joci", "Conceição do Almeida"},
		{"Emily", "Conceição do Almeida"},
		{"Hamilton", "Conceição do Almeida"},
		{"Sabrina", "Conceição do Almeida"},
		{"Lipe", "Juazeiro"},
		{"Torugo", "Itabuna"},
		{"Kai", "Conceição do Almeida"},
		{"Da pinta", "Juazeiro"},
		{"Will", "Conceição do Almeida"},
		{"Vi", "Conceição do Almeida"},
		{"Samires", "Conceição do Almeida"},
		{"Vitinho", "Conceição do Almeida"},
		{"Thai", "Conceição do Almeida"},
		{"Thiaguinho", "Itabuna"},
		{"Morena", "Conceição do Almeida"},
		{"Lana", "Conceição do Almeida"},
		{"Mucão", "Cruz das Almas"},
		{"Mel", "Canavieiras"},
		{"Mari Xavier", "Itabuna"},
		{"Lua", "Aracaju"},
		{"Ale", "Canavieiras"},
		{"Rê", "Aracaju"},
		{"Peu Amorim", "Jacobina"},
		{"Maria", "Jacobina"},
		{"Rodrigo", "Aracaju"},
		{"Nanda", "Itabuna"},
		{"Alef", "Conceição do Almeida"},
		{"Inha", "Conceição do Almeida"},
		{"Lua", "Conceição do Almeida"},
		{"Dri", "Conceição do Almeida"},
		{"Emylle", "Itabuna"},
		{"Ayssa Ribeiro", "Itabuna"},
		{"Leila", "Itabuna"},
		{"Dri", "Cruz das Almas"},
		{"Teté", "Conceição do Almeida"},
		{"Peri", "Conceição do Almeida"},
		{"Kaw", "Conceição do Almeida"},
		{"Kathy", "Conceição do Almeida"},
		{"Rato", "Conceição do Almeida"},
		{"Rol", "Conceição do Almeida"},
		{"Dan", "Conceição do Almeida"},
		{"Chele", "Aracaju"},
		{"Toinho", "Itabuna"},
		{"Clara", "Juazeiro"},
		{"Gabi", "Aracaju"},
		{"Formiga", "Canavieiras"},
		{"Malu", "Aracaju"},
		{"Ari", "Conceição do Almeida"},
		{"Nanda", "Conceição do Almeida"},
		{"Rick Marques", "Canavieiras"},
		{"Jan", "Conceição do Almeida"},
		{"Karen", "Conceição do Almeida"},
		{"Au", "Conceição do Almeida"},
		{"Manu", "Juazeiro"},
		{"Duda Souza", "Conceição do Almeida"},
		{"Kiti", "Conceição do Almeida"},
		{"Zé", "Conceição do Almeida"},
		{"Nathy", "Conceição do Almeida"},
		{"Gabi Rocha", "Aracaju"},
		{"Vivi", "Itabuna"},
		{"Gih", "Petrolina"},
		{"Mari", "Petrolina"},
		{"Biel", "Lauro de Freitas"},
		{"Dudinha", "Aracaju"},
		{"Ariel", "Aracaju"},
		{"Tia Helô", "Aracaju"},
		{"Re", "Aracaju"},
		{"Guilherme", "Aracaju"},
		{"Tio Emilio", "Aracaju"},
		{"Vivi", "Aracaju"},
		{"Lice", "Aracaju"},
		{"Isa Lauton", "Aracaju"},
		{"Raynight", "Aracaju"},
		{"Fernandinho", "Aracaju"},
		{"Momom", "Aracaju"},
		{"Lê Góes", "Aracaju"},
		{"Mama", "Aracaju"},
		{"Lu", "Aracaju"},
		{"Max", "Aracaju"},
		{"Jorge", "Aracaju"},
		{"Max", "Aracaju"},
		{"Bê Nascimento", "Aracaju"},
		{"Gabi Alves", "Aracaju"},
		{"Rabisco", "Jacobina"},
		{"Nai", "Cruz das Almas"},
		{"Vini", "Aracaju"},
		{"Taise", "Aracaju"},
		{"Mari", "Lauro de Freitas"},
		{"Garcez", "Aracaju"},
		{"Carioca", "Aracaju"},
		{"Nagy", "Aracaju"},
		{"Bia", "Cruz das Almas"},
		{"João Emannuel", "Juazeiro"},
		{"Mari", "Itabuna"},
		{"Ingrid", "Aracaju"},
		{"Lu", "Aracaju"},
		{"Tia Zu", "Aracaju"},
		{"Tia Cris", "Aracaju"},
		{"Rafa", "Petrolina"},
		{"Tio Kleber", "Aracaju"},
		{"Mari Sol", "Aracaju"},
		{"Gabriel", "Aracaju"},
		{"Bia", "Aracaju"},
		{"Bia Albernaz", "Cruz das Almas"},
		{"Leo", "Aracaju"},
		{"Teus", "Juazeiro"},
		{"Tio Ray", "Cruz das Almas"},
		{"Elisa", "Aracaju"},
		{"Lay", "Aracaju"},
		{"Girafa", "Cruz das Almas"},
		{"Ias", "Canavieiras"},
		{"Robazza", "Cruz das Almas"},
		{"Tio Joãozinho", "Aracaju"},
		{"Netinho", "Juazeiro"},
		{"Marcelo Hora", "Aracaju"},
		{"Bruninha", "Juazeiro"},
		{"Mi", "Aracaju"},
		{"Girassol", "Juazeiro"},
		{"Laura", "Itabuna"},
		{"Neri", "Juazeiro"},
		{"Dea", "Juazeiro"},
		{"Tonhão", "Aracaju"},
		{"Pipoca", "Lauro de Freitas"},
		{"Júlio", "Juazeiro"},
		{"Alice", "Aracaju"},
		{"Lu Carvalho", "Cruz das Almas"},
		{"Matheus", "Aracaju"},
		{"Bolinha", "Juazeiro"},
		{"Dai", "Juazeiro"},
		{"Guinho", "Itabuna"},
		{"Antonia", "Aracaju"},
		{"Tia Gina", "Aracaju"},
		{"Laurinha", "Aracaju"},
		{"Noah", "Cruz das Almas"},
		{"Chuchu", "Juazeiro"},
		{"Marina", "Aracaju"},
		{"Deza", "Cruz das Almas"},
		{"May", "Salvador"},
		{"Alaninha", "Ilhéus"},
		{"Agostinho", "Ilhéus"},
		{"Tia Alba", "Ilhéus"},
		{"Zezo", "Itabuna"},
		{"Alice", "Ilhéus"},
		{"Vitória", "Ilhéus"},
		{"Tio Tom", "Ilhéus"},
		{"Ayla", "Ilhéus"},
		{"Caquito", "Ilhéus"},
		{"Caio Rodrigues", "Ilhéus"},
		{"Carlinha", "Ilhéus"},
		{"Cathy", "Ilhéus"},
		{"Sara", "Ilhéus"},
		{"Tia Gal", "Ilhéus"},
		{"Ias", "Ilhéus"},
		{"Zé", "Ilhéus"},
		{"Isah", "Ilhéus"},
		{"Nino", "Ilhéus"},
		{"Jhulye", "Ilhéus"},
		{"Joanna", "Ilhéus"},
		{"Juli", "Ilhéus"},
		{"Juli", "Ilhéus"},
		{"Lavínia", "Ilhéus"},
		{"Livinha", "Ilhéus"},
		{"Liz", "Ilhéus"},
		{"Ló", "Ilhéus"},
		{"Felippe (Luizão)", "Ilhéus"},
		{"Marcus", "Ilhéus"},
		{"Meuri", "Ilhéus"},
		{"Loli", "Ilhéus"},
		{"Jota Pê", "Ilhéus"},
		{"Matheus", "Ilhéus"},
		{"Tia Nadir", "Ilhéus"},
		{"Longuinho", "Ilhéus"},
		{"Mano", "Aracaju"},
		{"Ray", "Ilhéus"},
		{"Valdir", "Ilhéus"},
		{"Júlio", "Ilhéus"},
		{"Vitor", "Ilhéus"},
		{"Vick", "Ilhéus"},
		{"Welder", "Ilhéus"},
		{"Flavinha", "Itabuna"},
		{"Gico", "Aracaju"},
		{"Bi", "Cruz das Almas"},
		{"Tio Marconi", "Aracaju"},
		{"Tia Rogéria", "Aracaju"},
		{"Cacá", "Cruz das Almas"},
		{"Tia Bela", "Cruz das Almas"},
		{"Jú", "Cruz das Almas"},
		{"Paulo Ivis", "Juazeiro"},
		{"Felipe", "Aracaju"},
		{"Lai", "Cruz das Almas"},
		{"Biel", "Cruz das Almas"},
		{"Luciana", "Juazeiro"},
		{"Míria", "Aracaju"},
		{"Bel", "Cruz das Almas"},
		{"Bia", "Aracaju"},
		{"Milla", "Juazeiro"},
		{"Marlandro", "Juazeiro"},
		{"Kaio", "Aracaju"},
		{"Artur", "Cruz das Almas"},
		{"Duda Torres", "Juazeiro"},
		{"Bia", "Itabuna"},
		{"Dorinha", "Aracaju"},
		{"JV Souza", "Juazeiro"},
		{"TuTu", "Aracaju"},
		{"Maria", "Aracaju"},
		{"Kauanzinho", "Itabuna"},
		{"Anninha", "Aracaju"},
		{"Nina", "Aracaju"},
		{"Lucas B", "Cruz das Almas"},
		{"Milly", "Juazeiro"},
		{"Matheus", "Juazeiro"},
		{"Teresa", "Aracaju"},
		{"Tia Mércia", "Aracaju"},
		{"Tio Vital", "Aracaju"},
		{"Laura", "Aracaju"},
		{"Renamm", "Juazeiro"},
		{"Anne", "Petrolina"},
		{"Diego", "Itabuna"},
		{"Gray", "Aracaju"},
		{"Júlia", "Aracaju"},
		{"Fê", "Petrolina"},
		{"DG", "Petrolina"},
		{"Tia Carminha", "Aracaju"},
		{"Aninha", "Juazeiro"},
		{"Mai", "Petrolina"},
		{"Duda", "Petrolina"},
		{"Arthurzinho", "Aracaju"},
		{"Sasa", "Aracaju"},
		{"Tia Veronica", "Aracaju"},
		{"Dudu", "Aracaju"},
		{"Rosi", "Aracaju"},
		{"Valverde", "Juazeiro"},
		{"Alvinho", "Juazeiro"},
		{"Toin", "Juazeiro"},
		{"Hugo", "Cruz das Almas"},
		{"Tia Su", "Conceição do Almeida"},
		{"Tia Júlia", "Aracaju"},
		{"Tio Sylvio", "Aracaju"},
		{"Mavi", "Cruz das Almas"},
		{"Maluzinha", "Cruz das Almas"},
		{"Mel", "Cruz das Almas"},
		{"Tio Carlos", "Aracaju"},
		{"Tia Dora", "Aracaju"},
		{"John", "Aracaju"},
		{"Mi", "Aracaju"},
		{"Bya", "Cruz das Almas"},
		{"Davi", "Aracaju"},
		{"Jefinho", "Cruz das Almas"},
		{"Carolzinha", "Aracaju"},
		{"Tio Eva", "Cruz das Almas"},
		{"Eugenio", "Aracaju"},
		{"Hugo", "Aracaju"},
		{"FF", "Aracaju"},
		{"Jay", "Aracaju"},
		{"Aninha", "Aracaju"},
		{"Jorginho", "Aracaju"},
		{"Belly", "Aracaju"},
		{"Bia Lemos", "Juazeiro"},
		{"Drica", "Aracaju"},
		{"May", "Aracaju"},
		{"Mago", "Aracaju"},
		{"Bella", "Aracaju"},
		{"Yara", "Aracaju"},
		{"AJ", "Cruz das Almas"},
		{"Dani P.", "Salvador"},
		{"MC", "Salvador"},
		{"Tia Juca", "Conceição do Almeida"},
		{"Ju", "Salvador"},
		{"Nana", "Salvador"},
		{"Carol Negredo", "Salvador"},
		{"Fabi Pitta", "Salvador"},
		{"Maju", "Salvador"},
		{"Luiza", "Salvador"},
		{"Nanda", "Salvador"},
		{"Nath", "Salvador"},
		{"Neto Raiz", "Aracaju"},
		{"Thai Lamarca", "Salvador"},
		{"Amanda Caroline", "Petrolina"},
		{"Leti Brito", "Salvador"},
		{"Theu", "Salvador"},
		{"Vilas", "Salvador"},
		{"Lai", "Salvador"},
		{"Rapha", "Salvador"},
		{"Luly", "Salvador"},
		{"Gi", "Salvador"},
		{"Kaio", "Salvador"},
		{"Gaby", "Salvador"},
		{"Paty", "Salvador"},
		{"Lipe", "Cruz das Almas"},
		{"Jiji", "Salvador"},
		{"Cacau", "Cruz das Almas"},
		{"Lari", "Salvador"},
		{"Mai", "Salvador"},
		{"Laurinha", "Salvador"},
		{"Amandinha C", "Salvador"},
		{"Cacá", "Salvador"},
		{"Lalai", "Salvador"},
		{"Nessa", "Salvador"},
		{"Renatinha", "Aracaju"},
		{"Glícia", "Aracaju"},
		{"Isa", "Salvador"},
		{"Wilson", "Aracaju"},
		{"Vitinho", "Salvador"},
		{"Rafa", "Salvador"},
		{"Ana Júlia", "Itabuna"},
		{"Mard", "Itabuna"},
		{"Gaby", "Itabuna"},
		{"Biscoito", "Itabuna"},
		{"Mari", "Itabuna"},
		{"Theu", "Itabuna"},
		{"Naty Barreto", "Itabuna"},
		{"João Ramos", "Itabuna"},
		{"João Preto", "Itabuna"},
		{"O nome do assunto", "Itabuna"},
		{"Biel", "Itabuna"},
		{"Galega", "Itabuna"},
		{"Paulinha", "Itabuna"},
		{"Xande", "Itabuna"},
		{"Luana", "Petrolina"},
		{"Tchuy", "Conceição do Almeida"},
		{"Nat Pedral", "Aracaju"},
		{"Ryan", "Petrolina"},
		{"Vic Coelho", "Petrolina"},
		{"Hen", "Petrolina"},
		{"Guigui", "Aracaju"},
	}
}
