CREATE TABLE IF NOT EXISTS posts (
    id INT PRIMARY KEY AUTO_INCREMENT,            
    user_id INT NOT NULL,                         
    title VARCHAR(255) NOT NULL,                  
    slug VARCHAR(255) NOT NULL,                   
    description TEXT DEFAULT NULL,                
    content TEXT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP DEFAULT NULL,
    FOREIGN KEY (user_id) REFERENCES users(id),
    UNIQUE (slug)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
