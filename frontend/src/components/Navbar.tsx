import React from 'react';
import { Box, Button } from '@chakra-ui/react';
import { useNavigate } from 'react-router-dom';

const Navbar: React.FC = () => {
  const navigate = useNavigate();
  return (
    <Box justifyContent="center" mt="4" mb="8">
      <Button onClick={() => navigate('/signup')} mr="4">
        Signup
      </Button>
      <Button onClick={() => navigate('/login')} mr="4">Login</Button>
      <Button onClick={() => navigate('/')} mr="4">Home</Button>
      <Button onClick={() => navigate('/todos')} mr="4">Todos</Button>
      <Button onClick={() => navigate('/logout')} mr="4">Logout</Button>
    </Box>
  );
};

export default Navbar;
