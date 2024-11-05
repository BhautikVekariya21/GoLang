// services.js contains functions for API requests
export const createBooking = async (bookingData) => {
    const response = await fetch('/api/bookings', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json',
        },
        body: JSON.stringify(bookingData),
    });
    return response.json();
};

export const createContact = async (contactData) => {
    const response = await fetch('/api/contact', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json',
        },
        body: JSON.stringify(contactData),
    });
    return response.ok;
};
