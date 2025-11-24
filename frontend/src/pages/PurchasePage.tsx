import { useAuth } from '../contexts/AuthContext';

export const PurchasePage = () => {
  const { user } = useAuth();

  return (
    <div style={{ padding: '20px' }}>
      <h1>Purchase Tickets</h1>
      <p>Welcome, {user?.email}!</p>
      <p>This is a protected page - only logged in users can see this.</p>

      <div style={{ marginTop: '20px' }}>
        <p>Ticket purchase form will be implemented in Phase 4.</p>
      </div>
    </div>
  );
};