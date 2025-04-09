-- Table des utilisateurs (inscription avec nom, prénom et mot de passe)
CREATE TABLE IF NOT EXISTS utilisateurs(
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    nom TEXT NOT NULL,
    prenom TEXT NOT NULL,
    phone_number TEXT NOT NULL UNIQUE,
    password_hash TEXT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Table des détails des clients (complété lors de la commande)
CREATE TABLE IF NOT EXISTS details_clients (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    client_id INTEGER NOT NULL,
    lieu_habitation TEXT NOT NULL,
    phone_number TEXT NOT NULL, -- Contact actif pour la prestation
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (client_id) REFERENCES utilisateurs(id)
);

-- Table des prestataires (travailleurs)
CREATE TABLE IF NOT EXISTS prestataires (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    name TEXT NOT NULL,
    service_type TEXT,
    phone_number TEXT NOT NULL UNIQUE,
    email TEXT,
    location TEXT,
    status TEXT CHECK(status IN ('available', 'unavailable')) DEFAULT 'available',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Table des services disponibles
CREATE TABLE IF NOT EXISTS services (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    service_name TEXT NOT NULL,
    description TEXT,
    is_urgent BOOLEAN DEFAULT FALSE,
    is_full_time BOOLEAN DEFAULT FALSE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Table des réservations de services
CREATE TABLE IF NOT EXISTS reservations (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    client_id INTEGER NOT NULL,
    prestataire_id INTEGER NOT NULL,
    service_id INTEGER NOT NULL,
    reservation_type TEXT CHECK(reservation_type IN ('urgent', 'appointment')) NOT NULL,
    appointment_time TIMESTAMP NOT NULL,
    status TEXT CHECK(status IN ('pending', 'confirmed', 'completed', 'cancelled')) DEFAULT 'pending',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (client_id) REFERENCES utilisateurs(id),
    FOREIGN KEY (prestataire_id) REFERENCES prestataires(id),
    FOREIGN KEY (service_id) REFERENCES services(id)
);

-- Table de chat (support client et contact avec prestataires)
CREATE TABLE IF NOT EXISTS chat_support (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    client_id INTEGER NOT NULL,
    prestataire_id INTEGER,
    message TEXT NOT NULL,
    sent_by TEXT CHECK(sent_by IN ('client', 'support', 'prestataire')) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (client_id) REFERENCES utilisateurs(id),
    FOREIGN KEY (prestataire_id) REFERENCES prestataires(id)
);

-- Table des paiements
CREATE TABLE IF NOT EXISTS paiements (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    reservation_id INTEGER NOT NULL,
    amount DECIMAL(10, 2) NOT NULL,
    payment_method TEXT NOT NULL,
    status TEXT CHECK(status IN ('pending', 'completed', 'failed')) DEFAULT 'pending',
    payment_date TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (reservation_id) REFERENCES reservations(id)
);
