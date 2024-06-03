import { BrowserRouter as Router, Routes, Route } from 'react-router-dom';
import { LoginPage } from "./pages/auth/LoginPage";
import { RegisterPage } from './pages/auth/RegisterPage';

export default function App() {
  return (
    <main>
        <Router>
            <Routes>
                <Route exact path='/' element={<LoginPage/>} />
                <Route exact path='/register' element={<RegisterPage/>} />
            </Routes>
        </Router>
    </main>
  );
}
