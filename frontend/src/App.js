import React from 'react';
import { BrowserRouter, Routes, Route} from 'react-router-dom';
import Login from './components/login/login'; // Importa el componente de Login
import MyComponent from './components/mycomponent';

function App() {
  const handleLogin = () => {
    // lógica de autenticación 
    console.log("Usuario autenticado");
  };

  return (
    <BrowserRouter>
      <Routes>
        <Route path="/" element={<MyComponent />} />
        <Route path="/login" element={<Login onLogin={handleLogin} />} /> 
      </Routes> 
    </BrowserRouter>
  );
}

export default App;