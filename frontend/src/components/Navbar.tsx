import { Link } from 'react-router-dom';
import { useAuth } from '../contexts/AuthContext';

export const Navbar = () => {
  const { user, isAuthenticated, logout } = useAuth();

  return (
    <nav style={{
      padding: '15px 20px',
      backgroundColor: '#333',
      color: 'white',
      display: 'flex',
      justifyContent: 'space-between',
      alignItems: 'center'
    }}>
      <div style={{ display: 'flex', gap: '20px' }}>
        <Link to="/events" style={{ color: 'white', textDecoration: 'none' }}>
          Events
        </Link>
        {isAuthenticated && (
          <Link to="/purchase" style={{ color: 'white', textDecoration: 'none' }}>
            Purchase
          </Link>
        )}
      </div>

      <div>
        {isAuthenticated ? (
          <div style={{ display: 'flex', gap: '15px', alignItems: 'center' }}>
            <span>Hi, {user?.email}</span>
            <button
              onClick={logout}
              style={{
                padding: '5px 15px',
                backgroundColor: '#dc3545',
                color: 'white',
                border: 'none',
                borderRadius: '4px',
                cursor: 'pointer'
              }}
            >
              Logout
            </button>
          </div>
        ) : (
          <Link to="/login">
            <button
              style={{
                padding: '5px 15px',
                backgroundColor: '#007bff',
                color: 'white',
                border: 'none',
                borderRadius: '4px',
                cursor: 'pointer'
              }}
            >
              Login
            </button>
          </Link>
        )}
      </div>
    </nav>
  );
};