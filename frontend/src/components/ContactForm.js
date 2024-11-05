import React, { useState } from 'react';

function ContactForm() {
    const [contactName, setContactName] = useState('');
    const [contactEmail, setContactEmail] = useState('');
    const [contactMessage, setContactMessage] = useState('');
    const [contactResponseMessage, setContactResponseMessage] = useState('');

    const handleContactSubmit = async (e) => {
        e.preventDefault();
        const response = await fetch('/api/contact', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
            },
            body: JSON.stringify({ name: contactName, email: contactEmail, message: contactMessage }),
        });

        if (response.ok) {
            setContactResponseMessage("Message sent successfully!");
        } else {
            setContactResponseMessage("Failed to send message.");
        }
    };

    return (
        <form onSubmit={handleContactSubmit}>
            <h2>Contact Us</h2>
            <input type="text" placeholder="Name" value={contactName} onChange={(e) => setContactName(e.target.value)} required />
            <input type="email" placeholder="Email" value={contactEmail} onChange={(e) => setContactEmail(e.target.value)} required />
            <textarea placeholder="Your Message" value={contactMessage} onChange={(e) => setContactMessage(e.target.value)} required />
            <button type="submit">Send Message</button>
            <div id="contactResponseMessage">{contactResponseMessage}</div>
        </form>
    );
}

export default ContactForm;
