import React from 'react';
import { Box } from '@chakra-ui/react';
import Navbar from './Navbar';

type LayoutProps = {
  children: React.ReactNode;
};

const Layout: React.FC<LayoutProps> = ({ children }) => {
  return (
    <Box>
      <Navbar />
      <Box w="100%" maxWidth="330px" mx="auto" pt="40px" pb="40px">
        {children}
      </Box>
    </Box>
  );
};

export default Layout;
