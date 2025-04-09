package controllers

import (
   
    "net/http"
    "html/template"
    "log"
    
)

func Home(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodGet {
        http.Error(w, "Méthode non autorisée", http.StatusMethodNotAllowed)
        return
    }


     // Charger les fichiers HTML : page principale, header, et footer
     tmpl, err := template.ParseFiles("frontend/templates/accueil.html", "frontend/templates/Header.html", "frontend/templates/footer.html")
     if err != nil {
         log.Println("Erreur lors du chargement des templates :", err)
         http.Error(w, "Erreur interne", http.StatusInternalServerError)
         return
     }

 
     // Exécuter le template et l'afficher
     err = tmpl.Execute(w, nil)
     if err != nil {
         log.Println("Erreur lors de l'exécution du template :", err)
         http.Error(w, "Erreur lors du rendu de la page", http.StatusInternalServerError)
    }
}

func LoginPage(w http.ResponseWriter, r *http.Request){
    log.Println("La fonction LoginPage a été appelée")
    if r.Method != http.MethodGet {
        http.Error(w, "Méthode non autorisée", http.StatusMethodNotAllowed)
        return
    }

   // Charger les fichiers HTML : page principale, header, et footer
   tmpl, err := template.ParseFiles("frontend/templates/login.html", "frontend/templates/Header.html", "frontend/templates/footer.html")
   if err != nil {
       log.Println("Erreur lors du chargement des templates :", err)
       http.Error(w, "Erreur interne", http.StatusInternalServerError)
       return
   }


   // Exécuter le template et l'afficher
   err = tmpl.Execute(w, nil)
   if err != nil {
       log.Println("Erreur lors de l'exécution du template :", err)
       http.Error(w, "Erreur lors du rendu de la page", http.StatusInternalServerError)
  }
}