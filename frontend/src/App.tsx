import { BrowserRouter as Router, Routes, Route } from 'react-router-dom'

function App() {
  return (
    <Router>
      <div className="app">
        <h1>Ticketing System</h1>
        <Routes>
          <Route path="/" element={<div>Home - Coming soon</div>} />
        </Routes>
      </div>
    </Router>
  )
}

export default App