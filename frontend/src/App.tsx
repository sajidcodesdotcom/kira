import { BrowserRouter, Route, Routes} from 'react-router-dom'
import Header from './components/header'
import HomePage from './pages/home_page'
import SignUpPage from './pages/signup_page'
import LoginPage from './pages/login_page'
import DashboardPage from './pages/dashboard_page'


function App() {
  return (
    <BrowserRouter>
    <Header />
    <Routes>
      <Route path="/" element={<HomePage  />} />
      <Route path="/signup" element={<SignUpPage />} />
      <Route path="/login" element={<LoginPage />} />
      <Route path="/dashboard" element={<DashboardPage />} />
    </Routes>
    </BrowserRouter>
  )
}

export default App
