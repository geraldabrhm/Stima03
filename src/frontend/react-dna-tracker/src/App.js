import './App.css';
import { Link, Route, Routes } from 'react-router-dom'
import Home from './pages/Home'
import InputDisease from './pages/InputDisease'
import DNATest from './pages/DNATest'
import DNATestResult from './pages/DNATestResult'

function App() {
  return (
    <div className="App">
      <nav>
        <li><Link to="/">Home</Link></li>
        <li><Link to="/input-disease">Input Disease</Link></li>
        <li><Link to="/dna-test">DNA Test</Link></li>
        <li><Link to="/dna-test-result">DNA Test Result</Link></li>
      </nav>
      
      <Routes>
        <Route path="/" element={<Home />} />
        <Route path="input-disease" element={<InputDisease />} />
        <Route path="dna-test" element={<DNATest />} />
        <Route path="dna-test-result" element={<DNATestResult />} />
      </Routes>
    </div>
  );
}

export default App;
