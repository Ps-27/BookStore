// src/components/BookList.js
import React, { useState, useEffect } from 'react';

const BookList = () => {
  const [books, setBooks] = useState([]);

  useEffect(() => {
    // Fetch the list of books from your Go API when the component mounts
    const fetchBooks = async () => {
      const response = await fetch('http://localhost:8080/books');
      if (response.ok) {
        const data = await response.json();
        setBooks(data);
      } else {
        console.error('Failed to fetch books');
      }
    };

    fetchBooks();
  }, []);
  return (
    <div>
      <h2>Book List</h2>
      <table>
        <thead>
          <tr>
            <th>ID</th>
            <th>Name</th>
            <th>Author</th>
            <th>Writer</th>
            <th>Publish Date</th>
          </tr>
        </thead>
        <tbody>
          {books.map((book) => (
            <tr key={book.id}>
              <td>{book.id}</td>
              <td>{book.name}</td>
              <td>{book.author}</td>
              <td>{book.writer}</td>
              <td>{book.publishDate}</td>
            </tr>
          ))} 
          </tbody>
      </table>
    </div>
  );
};

export default BookList;

  
