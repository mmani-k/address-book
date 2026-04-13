package com.example.addressbook.service;

import com.example.addressbook.model.Address;
import com.example.addressbook.repository.AddressRepository;
import org.springframework.stereotype.Service;

import java.util.List;

@Service
public class AddressService {

    private final AddressRepository addressRepository;

    public AddressService(AddressRepository addressRepository) {
        this.addressRepository = addressRepository;
    }

    public List<Address> getAll() {
        return addressRepository.findAll();
    }

    public Address getByName(String name) {
        return addressRepository.findById(name)
                .orElseThrow(() -> new IllegalArgumentException("Address not found for name: " + name));
    }

    public Address create(Address address) {
        if (addressRepository.existsById(address.getName())) {
            throw new IllegalArgumentException("Address already exists for name: " + address.getName());
        }
        return addressRepository.save(address);
    }

    public Address update(String name, Address updatedAddress) {
        Address existing = getByName(name);
        existing.setAddressLine1(updatedAddress.getAddressLine1());
        existing.setCity(updatedAddress.getCity());
        existing.setState(updatedAddress.getState());
        existing.setCountry(updatedAddress.getCountry());
        return addressRepository.save(existing);
    }

    public void delete(String name) {
        Address existing = getByName(name);
        addressRepository.delete(existing);
    }
}
