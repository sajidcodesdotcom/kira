import { BrowserRouter, Route, Routes} from 'react-router-dom'
import Header from './components/layout/header'
import HomePage from './pages/home_page'
import SignUpPage from './pages/signup_page'
import LoginPage from './pages/login_page'
import DashboardPage from './pages/dashboard_page'
import ProtectedRoute from './components/layout/protected_route'


function App() {
  return (
    <BrowserRouter>
    <Header />
    <Routes>
      <Route path="/" element={<HomePage  />} />
      <Route path="/signup" element={<SignUpPage />} />
      <Route path="/login" element={<LoginPage />} />
      <Route path="/dashboard" element={
        <ProtectedRoute>
        <DashboardPage />
        </ProtectedRoute>
        } />
    </Routes>
    </BrowserRouter>
  )
}

export default App
