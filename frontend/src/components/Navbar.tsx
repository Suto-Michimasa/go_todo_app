import React from 'react';
import { Box, Link } from '@chakra-ui/react';

const Navbar: React.FC = () => {
  return (
    <Box justifyContent="center" mt="4" mb="8">
      <Link href="/" mr="4">Top</Link>
      <Link href="/signup">Signup</Link>
    </Box>
  );
};

export default Navbar;
