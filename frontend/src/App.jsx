
import { useState, useEffect } from 'react';
import Login from './components/Login';
import Dashboard from './components/Dashboard';
import './App.css';

function App() {
  const [isAuthenticated, setIsAuthenticated] = useState(false);
  const [usuario, setUsuario] = useState(null);
  const [loading, setLoading] = useState(true);

  // Verificar si el usuario ya está autenticado al cargar la aplicación
  useEffect(() => {
    const authStatus = localStorage.getItem('isAuthenticated');
    const userData = localStorage.getItem('usuario');
    
    if (authStatus === 'true' && userData) {
      setIsAuthenticated(true);
      setUsuario(JSON.parse(userData));
    }
    setLoading(false);
  }, []);

  const handleLogin = (userData) => {
    setIsAuthenticated(true);
    setUsuario(userData);
  };

  const handleLogout = () => {
    setIsAuthenticated(false);
    setUsuario(null);
  };

  // Mostrar pantalla de carga mientras se verifica la autenticación
  if (loading) {
    return (
      <div className="min-h-screen flex items-center justify-center">
        <div className="animate-spin rounded-full h-32 w-32 border-b-2 border-indigo-600"></div>
      </div>
    );
  }

  return (
    <>
      {isAuthenticated ? (
        <Dashboard usuario={usuario} onLogout={handleLogout} />
      ) : (
        <Login onLogin={handleLogin} />
      )}
    </>
  );
}

export default App;
