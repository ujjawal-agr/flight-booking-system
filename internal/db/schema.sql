-- Enable UUID generation
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

-- ENUM definitions
CREATE TYPE flight_status AS ENUM ('available', 'cancelled', 'fullyBooked');
CREATE TYPE seat_type AS ENUM ('window', 'middle', 'aisle');
CREATE TYPE seat_status AS ENUM ('available', 'booked');
CREATE TYPE gender_enum AS ENUM ('male', 'female', 'other');

-- Table: Flight
CREATE TABLE flights (
                         flight_id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
                         flight_code VARCHAR(20) UNIQUE NOT NULL ,
                         source VARCHAR(50) NOT NULL,
                         destination VARCHAR(50) NOT NULL,
                         company VARCHAR(50) NOT NULL,
                         flight_status flight_status NOT NULL DEFAULT 'available',
                         date TIMESTAMP NOT NULL,
                         created_on TIMESTAMP DEFAULT NOW(),
                         updated_on TIMESTAMP DEFAULT NOW(),
                         created_by VARCHAR(50),
                         updated_by VARCHAR(50)
);

-- Table: Seat
CREATE TABLE seats (
                       seat_id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
                       flight_id UUID REFERENCES flights(flight_id) ON DELETE CASCADE,
                       seat_no INT NOT NULL,
                       seat_type seat_type NOT NULL,
                       seat_status seat_status NOT NULL DEFAULT 'available',
                       created_on TIMESTAMP DEFAULT NOW(),
                       updated_on TIMESTAMP DEFAULT NOW(),
                       created_by VARCHAR(50),
                       updated_by VARCHAR(50)
);

-- Table: Pricing
CREATE TABLE pricing (
                         price_id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
                         flight_id UUID REFERENCES flights(flight_id) ON DELETE CASCADE,
                         seat_type seat_type NOT NULL,
                         price INT NOT NULL,
                         created_on TIMESTAMP DEFAULT NOW(),
                         updated_on TIMESTAMP DEFAULT NOW(),
                         created_by VARCHAR(50),
                         updated_by VARCHAR(50)
);

-- Table: Booking Details
CREATE TABLE bookings (
                          booking_id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
                          flight_id UUID REFERENCES flights(flight_id) ON DELETE CASCADE,
                          customer_name VARCHAR(100) NOT NULL,
                          customer_contact VARCHAR(20),
                          created_on TIMESTAMP DEFAULT NOW(),
                          updated_on TIMESTAMP DEFAULT NOW(),
                          created_by VARCHAR(50),
                          updated_by VARCHAR(50)
);

-- Table: Booking-Seat Mapping
CREATE TABLE booking_seat_mapping (
                                      map_id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
                                      booking_id UUID REFERENCES bookings(booking_id) ON DELETE CASCADE,
                                      seat_id UUID REFERENCES seats(seat_id) ON DELETE CASCADE,
                                      name VARCHAR(100) NOT NULL,
                                      age INT NOT NULL,
                                      gender gender_enum NOT NULL,
                                      created_on TIMESTAMP DEFAULT NOW(),
                                      updated_on TIMESTAMP DEFAULT NOW(),
                                      created_by VARCHAR(50),
                                      updated_by VARCHAR(50)
);


