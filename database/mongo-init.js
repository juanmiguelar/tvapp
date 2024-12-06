// Read environment variables
const user = process.env.MONGO_INITDB_ROOT_USERNAME;
const password = process.env.MONGO_INITDB_ROOT_PASSWORD;
const database = process.env.MONGO_DATABASE;

// Connect to MongoDB using credentials from environment variables
db = connect(`mongodb://${user}:${password}@localhost:27017/admin`);

// Create or connect to the specified database
db = db.getSiblingDB(database);

// Insert initial data with authors
db.news.insertMany([
    {
        title: "Breaking News: AI Advances",
        content: "AI technology is advancing at an unprecedented rate.",
        author: {
            name: "Alice Johnson",
            email: "alice.johnson@example.com"
        }
    },
    {
        title: "Tech News: New Programming Language Released",
        content: "A new programming language, QuantumLang, has been released for quantum computing enthusiasts.",
        author: {
            name: "Bob Smith",
            email: "bob.smith@example.com"
        }
    },
    {
        title: "Health Update: Benefits of Daily Exercise",
        content: "Studies show that 30 minutes of daily exercise can improve mental health.",
        author: {
            name: "Charlie Lee",
            email: "charlie.lee@example.com"
        }
    },
    {
        title: "Economy: Stock Market Hits Record High",
        content: "The stock market reached an all-time high today, driven by tech stocks.",
        author: {
            name: "Diana Roberts",
            email: "diana.roberts@example.com"
        }
    }
]);

db = db.getSiblingDB("admin"); // Switch to the admin database

db.createUser({
    user: user,
    pwd: password, // Use the same password as in your .env
    roles: [
        { role: "userAdminAnyDatabase", db: "admin" },
        { role: "readWriteAnyDatabase", db: "admin" }
    ]
});
