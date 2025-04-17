import React, { useEffect, useState } from 'react';

function App() {
  const [users, setUsers] = useState([]);
  const API_URL = process.env.REACT_APP_API_URL || "http://localhost:8080";

  useEffect(() => {
    fetch(`${API_URL}/users`)
      .then((res) => res.text())
      .then((text) => {
        const lines = text.trim().split('\n').slice(1); // skip "Users:"
        const parsedUsers = lines.map(line => {
          const [id, name] = line.split(': ');
          return { id, name };
        });
        setUsers(parsedUsers);
      });
  }, []);

  return (
    <div style={{ padding: '2rem' }}>
      <h1>Lista de Usuarios</h1>
      <ul>
        {users.map(user => (
          <li key={user.id}>{user.name} (ID: {user.id})</li>
        ))}
      </ul>
    </div>
  );
}

export default App;
