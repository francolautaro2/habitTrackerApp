import React, { useState } from 'react';
import axios from 'axios';
import {jwtDecode} from 'jwt-decode';
import './login.css'; // Importar el archivo de estilos CSS desde la carpeta styles

const Login = () => {
  const [formData, setFormData] = useState({
    username_or_email: '',
    password: ''
  });

  const handleChange = (e) => {
    setFormData({ ...formData, [e.target.name]: e.target.value });
  };

  const handleSubmit = async (e) => {
    e.preventDefault();
    try {
      const response = await axios.post('http://localhost:8080/api/login', formData, {
        headers: {
          'Content-Type': 'application/json'
        }
      });
      const token = response.data.token;
      const decodedToken = jwtDecode(token);
      console.log('Informacion del token: ', decodedToken); 
      // Aquí puedes manejar la respuesta del servidor, como guardar el token JWT en el almacenamiento local
    } catch (error) {
      console.error('Error de inicio de sesión:', error);
    }
  };

  return (
    <div className="login-container"> {/* Aplicar clase CSS para contenedor principal */}
      <form className="login-form" onSubmit={handleSubmit}> {/* Aplicar clase CSS para formulario */}
        <input className="login-input" type="text" name="username_or_email" placeholder="Usuario o email" value={formData.username_or_email} onChange={handleChange} />
        <input className="login-input" type="password" name="password" placeholder="Contraseña" value={formData.password} onChange={handleChange} />
        <button className="login-button" type="submit">Iniciar sesión</button>
        <a href="/register" className="register-link">¿No tienes una cuenta? Regístrate aquí</a> {/* Añadir enlace de registro */}
      </form>
    </div>
  );
};

export default Login;