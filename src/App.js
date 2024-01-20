import 'bootswatch/dist/darkly/bootstrap.min.css';
import './App.css';
import Dump from './components/Dump';
import Header from './components/Header';

function App() {
  return (
    <div className="App">
      <Header />
      <Dump />
      {/* <header className="App-header">
        <img src={logo} className="App-logo" alt="logo" />
        <p>
          Edit <code>src/App.js</code> and save to reload.
        </p>
        <a
          className="App-link"
          href="https://reactjs.org"
          target="_blank"
          rel="noopener noreferrer"
        >
          Learn React
        </a>
      </header> */}
    </div>
  );
}

export default App;
