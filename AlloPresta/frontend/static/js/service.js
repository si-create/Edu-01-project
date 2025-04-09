 // Fonction pour afficher ou masquer les boutons en fonction du choix
 function updateButtons(showEmergency, showValidate) {
    const emergencyButton = document.getElementById('emergencyButton');
    const appointmentButton = document.getElementById('appointmentButton');
    const validateButton = document.getElementById('validateButton');
    const contractDetails = document.getElementById('contractDetails');

    // Gérer les boutons
    if (showEmergency) {
        emergencyButton.classList.add('visible');
        emergencyButton.classList.remove('hidden');
        appointmentButton.classList.add('visible');
        appointmentButton.classList.remove('hidden');
        validateButton.classList.add('hidden');
        validateButton.classList.remove('visible');
    } else if (showValidate) {
        validateButton.classList.add('visible');
        validateButton.classList.remove('hidden');
        emergencyButton.classList.add('hidden');
        emergencyButton.classList.remove('visible');
        appointmentButton.classList.add('hidden');
        appointmentButton.classList.remove('visible');
    } else {
        emergencyButton.classList.add('hidden');
        emergencyButton.classList.remove('visible');
        appointmentButton.classList.add('hidden');
        appointmentButton.classList.remove('visible');
        validateButton.classList.add('hidden');
        validateButton.classList.remove('visible');
    }

    // Gérer l'affichage des détails pour le contrat mensuel
    if (showValidate) {
        contractDetails.classList.add('visible');
        contractDetails.classList.remove('hidden');
    } else {
        contractDetails.classList.add('hidden');
        contractDetails.classList.remove('visible');
    }
}

function loadHeader() {
    const container = document.getElementById('header-container');
    container.innerHTML = ''; // Vide le header avant d'ajouter le nouveau contenu

    fetch('./Header.html')
        .then(response => response.text())
        .then(data => {
            container.innerHTML = data;
        })
        .catch(error => console.error('Error loading header:', error));
}

function loadFooter() {
    const container = document.getElementById("footer-container");
    container.innerHTML = ''; // Vide le footer avant d'ajouter le nouveau contenu

    fetch('./footer.html')
        .then(response => response.text())
        .then(data => {
            container.innerHTML = data;
        })
        .catch(error => console.error('Error loading footer:', error));
}


// Appeler la fonction pour charger le header et le footer dès que le DOM est prêt
document.addEventListener('DOMContentLoaded', function() {
    loadHeader();
    loadFooter();
});

// Fonction pour afficher/masquer les formulaires de connexion et d'inscription
function showForm(formType) {
    // Cacher les deux formulaires par défaut
    document.getElementById('connexion-form').classList.add('hidden');
    document.getElementById('inscription-form').classList.add('hidden');

    // Afficher le formulaire correspondant
    if (formType === 'connexion') {
        document.getElementById('connexion-form').classList.remove('hidden');
        document.getElementById('connexion-tab').classList.add('border-b-2', 'border-[#924C10]');
        document.getElementById('inscription-tab').classList.remove('border-b-2', 'border-[#924C10]');
    } else {
        document.getElementById('inscription-form').classList.remove('hidden');
        document.getElementById('inscription-tab').classList.add('border-b-2', 'border-[#924C10]');
        document.getElementById('connexion-tab').classList.remove('border-b-2', 'border-[#924C10]');
    }
}

// Initialisation - Le formulaire de connexion s'affiche par défaut au premier chargement
window.onload = function() {
    showForm('connexion');
};

// Script pour afficher les détails du service dynamiquement
document.addEventListener("DOMContentLoaded", function () {
    const params = new URLSearchParams(window.location.search);
    const service = params.get("service");

    const details = {
        menage: {
            title: "Service de Ménage",
            description: "Nous vous proposons un service de ménage complet pour un intérieur propre et accueillant.",
            image: "../static/image/services/femme-de-menage.png"
        },
        lessive: {
            title: "Service de Lessive",
            description: "Prenez soin de vos vêtements sans effort avec notre service de lessive professionnel.",
            image: "../static/image/services/la-lessive.png"
        },
        nounou: {
            title: "Service de Nounou",
            description: "Confiez vos enfants à des nounous qualifiées et expérimentées pour votre tranquillité d'esprit.",
            image: "../static/image/services/nounou.png"
        },
        cuisiniere: {
            title: "Service de Cuisinière",
            description: "Dégustez des repas faits maison grâce à nos cuisinières talentueuses.",
            image: "../static/image/services/nourriture-chinoise.png"
        }
    };

    if (service && details[service]) {
        document.getElementById("service-title").innerText = details[service].title;
        document.getElementById("service-description").innerText = details[service].description;
        document.getElementById("service-image").src = details[service].image;
    } else {
        document.getElementById("service-title").innerText = "Service non trouvé";
        document.getElementById("service-description").innerText = "Le service demandé n'existe pas.";
    }

    // Gestion de l'affichage du champ "Nombre de douches"
    const doucheCheckbox = document.getElementById("douche");
    const doucheNumberField = document.getElementById("nombre-douches-container");

    doucheCheckbox.addEventListener("change", function () {
        if (this.checked) {
            doucheNumberField.classList.remove("hidden");
        } else {
            doucheNumberField.classList.add("hidden");
        }
    });
});

// Fonction pour gérer la soumission du formulaire
function handleSubmit(event) {
    event.preventDefault();
    const habits = document.getElementById("habits").value;
    const draps = document.getElementById("draps").value;
    const serviettes = document.getElementById("serviettes").value;
    const jeans = document.getElementById("jeans").value;
    const douche = document.getElementById("douche").checked;
    const nombreDouches = douche ? document.getElementById("nombre-douches").value : "Non spécifié";

    alert(`Détails envoyés : 
- Nombre d'habits : ${habits}
- Nombre de draps : ${draps}
- Nombre de serviettes : ${serviettes}
- Nombre de jeans : ${jeans}
- Douche : ${douche ? "Oui" : "Non"}
- Nombre de douches : ${nombreDouches}`);
}

// Fonction pour afficher ou masquer le champ de la date en fonction de l'option sélectionnée
function toggleDateField(isReservation) {
    const reservationDateField = document.getElementById('reservationDateField');
    if (isReservation) {
        reservationDateField.classList.remove('hidden');
    } else {
        reservationDateField.classList.add('hidden');
    }
}

// Gestion de l'envoi du formulaire
document.getElementById('orderForm').addEventListener('submit', function(event) {
    event.preventDefault();
    
    // Récupérer l'option choisie
    const orderType = document.querySelector('input[name="orderType"]:checked');
    const reservationDate = document.getElementById('reservationDate').value;

    if (orderType) {
        if (orderType.id === 'reservation' && !reservationDate) {
            alert('Veuillez choisir une date de prestation.');
            return;
        }
        // Afficher un message de confirmation (ou rediriger vers une autre page)
        alert('Votre commande a bien été enregistrée !');
    } else {
        alert('Veuillez sélectionner une option (Urgente ou Réservation).');
    }
});

