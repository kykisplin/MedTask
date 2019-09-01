package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type AllMedicines struct {
	Categ Category
	Medic []Medicines
}

type Medicines struct {
	ID      string `json:"id"`
	NameMed string `json:"name"`
	IDcat   string `json: ID category`
}

type Category struct {
	ID      string `json: "id"`
	NameCat string `json: "Name"`
	//Medic   []Medicines
}

var med = []Medicines{}
var cat = []Category{}

//Получение категорий и медикаментов в них
func getAll(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	//params := mux.Vars(r)

	var allMed = []AllMedicines{}

	for _, ItemCat := range cat {
		newlist := []Medicines{}
		for _, ItemMed := range med {
			if ItemMed.IDcat == ItemCat.ID {
				newlist = append(newlist, ItemMed)
			}
		}

		allMed = append(allMed, AllMedicines{ItemCat, newlist})
	}

	json.NewEncoder(w).Encode(allMed)
}

//////////////// Работа с медикаментами////////////////////////////////////

// Получение всех медикаментов
func getMed(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(med)
}

//Получение информации о медикаменте
func getMedicines(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	for _, item := range med {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}

	json.NewEncoder(w).Encode(&Medicines{})
}

//Создание медикамента
func createMedicines(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var medicin Medicines

	_ = json.NewDecoder(r.Body).Decode(&medicin)
	med = append(med, medicin)

	json.NewEncoder(w).Encode(medicin)
}

//редактирование медикаментов
func updateMedicines(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range med {
		if item.ID == params["id"] {
			med = append(med[:index], med[index+1:]...)
			var TMed Medicines
			_ = json.NewDecoder(r.Body).Decode(&TMed)
			TMed.ID = params["id"]
			med = append(med, TMed)
			json.NewEncoder(w).Encode(TMed)
			return
		}
	}
	json.NewEncoder(w).Encode(med)
}

//удаление медикаментов
func deleteMedicines(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range med {
		if item.ID == params["id"] {
			med = append(med[:index], med[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(med)
}

//////////////// Работа с категориями////////////////////////////////////

//Получение всех категорий
func getCat(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(cat)
}

//Получение информации о категории
func getCategory(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _, item := range cat {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}

	json.NewEncoder(w).Encode(&Category{})
}

//Создание Категории
func createCategory(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var categ Category

	_ = json.NewDecoder(r.Body).Decode(&categ)
	cat = append(cat, categ)

	json.NewEncoder(w).Encode(categ)
}

//редактирование категории
func updateCategory(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range cat {
		if item.ID == params["id"] {
			cat = append(cat[:index], cat[index+1:]...)
			var TCat Category
			_ = json.NewDecoder(r.Body).Decode(&TCat)
			TCat.ID = params["id"]
			cat = append(cat, TCat)
			json.NewEncoder(w).Encode(TCat)
			return
		}
	}
	json.NewEncoder(w).Encode(cat)
}

//удаление категории
func deleteCategory(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range cat {
		if item.ID == params["id"] {
			cat = append(cat[:index], cat[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(cat)
}

func main() {

	r := mux.NewRouter()
	med = append(med, Medicines{ID: "0", NameMed: "Парацетамол", IDcat: "1"})
	med = append(med, Medicines{ID: "1", NameMed: "Аскофен", IDcat: "0"})
	cat = append(cat, Category{ID: "0", NameCat: "Жаропонижающие"})
	cat = append(cat, Category{ID: "1", NameCat: "Анаболики"})

	//Получение всей информации
	r.HandleFunc("/all", getAll).Methods("GET")

	//работа с категориями
	r.HandleFunc("/cat", getCat).Methods("GET")
	r.HandleFunc("/cat/{id}", getCategory).Methods("GET")
	r.HandleFunc("/cat", createCategory).Methods("POST")
	r.HandleFunc("/cat/{id}", updateCategory).Methods("PUT")
	r.HandleFunc("/cat/{id}", deleteCategory).Methods("DELETE")

	//работа с медикаментами
	r.HandleFunc("/med", getMed).Methods("GET")
	r.HandleFunc("/med/{id}", getMedicines).Methods("GET")
	r.HandleFunc("/med", createMedicines).Methods("POST")
	r.HandleFunc("/med/{id}", updateMedicines).Methods("PUT")
	r.HandleFunc("/med/{id}", deleteMedicines).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8080", r))

}
