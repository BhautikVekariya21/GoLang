import React, { useState } from 'react';

function BookingForm() {
    const [name, setName] = useState('');
    const [email, setEmail] = useState('');
    const [eventDate, setEventDate] = useState('');
    const [guestCount, setGuestCount] = useState('');
    const [decorationCharges, setDecorationCharges] = useState(0);
    const [lightBill, setLightBill] = useState(0);
    const [responseMessage, setResponseMessage] = useState('');

    const handleSubmit = async (e) => {
        e.preventDefault();
        const response = await fetch('/api/bookings', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
            },
            body: JSON.stringify({ name, email, eventDate, guestCount, decorationCharges, lightBill }),
        });

        const data = await response.json();
        setResponseMessage(data.message + (data.totalCost ? ` Total Cost: ₹${data.totalCost}` : ''));
    };

    return (
        <form onSubmit={handleSubmit}>
            <h2>Book Your Event</h2>
            <input type="text" placeholder="Name" value={name} onChange={(e) => setName(e.target.value)} required />
            <input type="email" placeholder="Email" value={email} onChange={(e) => setEmail(e.target.value)} required />
            <input type="date" value={eventDate} onChange={(e) => setEventDate(e.target.value)} required />
            <input type="number" placeholder="Number of Guests" value={guestCount} onChange={(e) => setGuestCount(e.target.value)} required />
            <input type="number" placeholder="Decoration Charges" value={decorationCharges} onChange={(e) => setDecorationCharges(e.target.value)} />
            <input type="number" placeholder="Light Bill" value={lightBill} onChange={(e) => setLightBill(e.target.value)} />
            <button type="submit">Book Now</button>
            <div id="responseMessage">{responseMessage}</div>
        </form>
    );
}

export default BookingForm;
