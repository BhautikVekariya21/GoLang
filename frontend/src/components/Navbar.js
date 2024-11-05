import React from 'react';
import { Link } from 'react-router-dom';

function Navbar() {
    return (
        <nav style={styles.nav}>
            <Link to="/" style={styles.link}>Home</Link>
            <Link to="/services" style={styles.link}>Services</Link>
            <Link to="/contact" style={styles.link}>Contact Us</Link>
        </nav>
    );
}

const styles = {
    nav: {
        display: 'flex',
        justifyContent: 'center', // Center the navbar items
        backgroundColor: '#D8BFD8', // Lavender shade for navbar
        padding: '10px 0', // Padding for the navbar
        borderRadius: '8px', // Rounded corners for the navbar
        boxShadow: '0 2px 5px rgba(0, 0, 0, 0.1)', // Subtle shadow
    },
    link: {
        margin: '0 15px',
        textDecoration: 'none',
        color: 'darkslateblue', // Darker link color for contrast
        fontWeight: 'bold', // Make the links bold
    }
};

export default Navbar;
