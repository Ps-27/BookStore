// src/App.js
import React from 'react';
import DefaultBooks from './components/DefaultBooks';
import HighestDemandingBooks from './components/HighestDemandingBooks';
import Welcome from './components/Welcome';


function App() {
  const [books, setBooks] = useState([]);
  const [searchResults, setSearchResults] = useState([]);

  useEffect(() => {
    // Fetch all books initially (or you can fetch on component mount)
    fetchBooks();
  }, []);

  const fetchBooks = () => {
    axios.get('http://your-api-url/books')
      .then((response) => {
        setBooks(response.data);
      })
      .catch((error) => {
        console.error('Error fetching books:', error);
      });
  };
  const handleSearch = (query) => {
    // Send a request to the backend to search for books based on the query
    axios.get(`http://your-api-url/search?query=${query}`)
      .then((response) => {
        setSearchResults(response.data);
      })
      .catch((error) => {
        console.error('Error searching for books:', error);
      });
  };

  return (
    <div className="App">
      <h1>Bookstore</h1>
      <h1>Bookstore</h1>
      <Welcome />
      <DefaultBooks />
      <HighestDemandingBooks />
      <Search onSearch={handleSearch} />
      {searchResults.length > 0 ? (
        <BookList books={searchResults} />
      ) : (
        <BookList books={books} />
      )}
    </div>
 );
}

export default App;

