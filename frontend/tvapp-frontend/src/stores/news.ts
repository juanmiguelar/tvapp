import { defineStore } from "pinia";
import axios from "axios";

export const useNewsStore = defineStore("news", {
  state: () => ({
    newsList: [] as Array<{
      id: string;
      title: string;
      content: string;
      author: { name: string; email: string };
    }>,
    loading: false,
    error: null as string | null,
  }),
  actions: {
    async fetchNews() {
      try {
        this.loading = true;
        const response = await axios.post("/query", {
          query: `query { getNews { id title content author { name email } } }`,
        });
        this.newsList = response.data.data.getNews;
      } catch (err: any) {
        this.error = err.message || "Failed to fetch news.";
      } finally {
        this.loading = false;
      }
    },
    async createNews(newNews: {
      title: string;
      content: string;
      authorName: string;
      authorEmail: string;
    }) {
      try {
        const response = await axios.post("/query", {
          query: `mutation {
            createNews(
              title: "${newNews.title}",
              content: "${newNews.content}",
              authorName: "${newNews.authorName}",
              authorEmail: "${newNews.authorEmail}"
            ) {
              id
              title
              content
              author {
                name
                email
              }
            }
          }`,
        });
        this.newsList.push(response.data.data.createNews); // Add the newly created news to the list
      } catch (err: any) {
        throw new Error(err.message || "Failed to create news.");
      }
    },
    async deleteNews(id: string) {
      try {
        const response = await axios.post("/query", {
          query: `mutation{deleteNews(id:"${id}")}`,
        });
        // Check if the backend indicates success
        const result = response.data?.data?.deleteNews;
        if (result) {
          this.newsList = this.newsList.filter((news) => news.id !== id); // Update the local state
        } else {
          this.newsList = this.newsList.filter((news) => news.id !== id);
          throw new Error("Backend failed to delete the news.");
        }
      } catch (err: any) {
        console.error("Error deleting news:", err);
        throw new Error(err.message || "Failed to delete news.");
      }
    }
  },
});
