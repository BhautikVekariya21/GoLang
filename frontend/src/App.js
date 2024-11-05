import React from 'react';
import { BrowserRouter as Router, Route, Routes } from 'react-router-dom';
import BookingForm from './components/BookingForm';
import ContactForm from './components/ContactForm';
import Navbar from './components/Navbar';

function App() {
    return (
        <Router>
            <Navbar />
            <Routes>
                <Route path="/" element={<BookingForm />} />
                <Route path="/contact" element={<ContactForm />} />
                <Route path="/services" element={<h2>Our Services</h2>} />
            </Routes>
        </Router>
    );
}

export default App;
