import { useAuth } from './context/AuthContext';


import './App.css'

function App() {
  const {  logout } = useAuth();

  return (
    <div>
        <>
          <p>Bem-vindo, usuário!</p>
          <button onClick={logout}>Logout</button>
        </>
 
    </div>
  );
}

export default App
