import { BrowserRouter, Route, Routes} from 'react-router-dom'
import Header from './components/header'
import HomePage from './pages/home-page'
import SignUpPage from './pages/signup-page'


function App() {
  return (
    <BrowserRouter>
    <Header />
    <Routes>
      <Route path="/signup" element={<SignUpPage />} />
      <Route path="/" element={<HomePage  />} />
    </Routes>
    </BrowserRouter>
  )
}

export default App
