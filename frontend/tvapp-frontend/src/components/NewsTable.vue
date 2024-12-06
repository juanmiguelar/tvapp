<template>
  <div class="news-table">
    <h1>News List</h1>
    <button @click="toggleForm" class="create-button">
      {{ showForm ? "Cancel" : "Add News" }}
    </button>

    <!-- Add News Form -->
    <div v-if="showForm" class="news-form">
      <h2>Add News</h2>
      <form @submit.prevent="addNews">
        <div>
          <label for="title">Title:</label>
          <input v-model="form.title" id="title" type="text" required />
        </div>
        <div>
          <label for="content">Content:</label>
          <textarea v-model="form.content" id="content" required></textarea>
        </div>
        <div>
          <label for="authorName">Author Name:</label>
          <input v-model="form.authorName" id="authorName" type="text" required />
        </div>
        <div>
          <label for="authorEmail">Author Email:</label>
          <input v-model="form.authorEmail" id="authorEmail" type="email" required />
        </div>
        <button type="submit">Submit</button>
      </form>
    </div>

    <!-- Display Loading/Error -->
    <div v-if="newsStore.loading" class="loading">Loading...</div>
    <div v-if="newsStore.error" class="error">{{ newsStore.error }}</div>
    <div v-if="!newsStore.loading && !newsStore.error && newsStore.newsList.length === 0" class="empty">
      No news found.
    </div>

    <!-- News Table -->
    <table v-if="!newsStore.loading && !newsStore.error && newsStore.newsList.length > 0">
      <thead>
        <tr>
          <th>Title</th>
          <th>Content</th>
          <th>Author</th>
          <th>Email</th>
          <th>Actions</th>
        </tr>
      </thead>
      <tbody>
        <tr v-for="news in newsStore.newsList" :key="news.id">
          <td>{{ news.title }}</td>
          <td>{{ news.content }}</td>
          <td>{{ news.author.name }}</td>
          <td>{{ news.author.email }}</td>
          <td>
            <button @click="deleteNews(news.id)" class="delete-button">Delete</button>
          </td>
        </tr>
      </tbody>
    </table>
  </div>
</template>

<script lang="ts">
import { defineComponent } from "vue";
import { useNewsStore } from "../stores/news";

export default defineComponent({
  name: "NewsTable",
  data() {
    return {
      showForm: false,
      form: {
        title: "",
        content: "",
        authorName: "",
        authorEmail: "",
      },
    };
  },
  computed: {
    newsStore() {
      return useNewsStore(); // Access the Pinia store
    },
  },
  methods: {
    toggleForm() {
      this.showForm = !this.showForm; // Toggle form visibility
    },
    async addNews() {
      try {
        await this.newsStore.createNews(this.form); // Add news via store method
        this.resetForm(); // Reset form fields
      } catch (error) {
        console.error("Failed to add news:", error);
      }
    },
    async deleteNews(id: string) {
      try {
        await this.newsStore.deleteNews(id); // Delete news via store method
      } catch (error) {
        console.error("Failed to delete news:", error);
      }
    },
    resetForm() {
      this.form = { title: "", content: "", authorName: "", authorEmail: "" }; // Clear form fields
      this.showForm = false; // Hide form
    },
  },
  mounted() {
    this.newsStore.fetchNews(); // Fetch news when the component is mounted
  },
});
</script>

<style>
.news-table {
  padding: 20px;
  font-family: Arial, sans-serif;
}

.create-button {
  margin-bottom: 20px;
  padding: 10px 20px;
  font-size: 1em;
  cursor: pointer;
  background-color: #4caf50;
  color: white;
  border: none;
  border-radius: 4px;
}

.news-form {
  margin-bottom: 20px;
  padding: 20px;
  border: 1px solid #ddd;
  background-color: #f9f9f9;
}

.news-form h2 {
  margin-bottom: 10px;
}

.news-form form > div {
  margin-bottom: 10px;
}

.news-form label {
  display: block;
  font-weight: bold;
}

.news-form input,
.news-form textarea {
  width: 100%;
  padding: 8px;
  margin-top: 4px;
  border: 1px solid #ddd;
  border-radius: 4px;
}

.news-form button {
  padding: 10px 20px;
  font-size: 1em;
  cursor: pointer;
  background-color: #007bff;
  color: white;
  border: none;
  border-radius: 4px;
}

.loading,
.error,
.empty {
  margin: 20px 0;
  font-size: 1.2em;
  color: #666;
}

table {
  width: 100%;
  border-collapse: collapse;
  margin-top: 20px;
}

th,
td {
  border: 1px solid #ddd;
  padding: 8px;
}

th {
  background-color: #f4f4f4;
  text-align: left;
}

tbody tr:hover {
  background-color: #f9f9f9;
}

.error {
  color: red;
}

.empty {
  color: #888;
}

.delete-button {
  padding: 5px 10px;
  background-color: #e74c3c;
  color: white;
  border: none;
  border-radius: 4px;
  cursor: pointer;
}

.delete-button:hover {
  background-color: #c0392b;
}
</style>
