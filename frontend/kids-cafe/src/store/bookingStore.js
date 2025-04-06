import { defineStore } from "pinia";

export const useBookingStore = defineStore("booking", {
  state: () => ({
    roomID: "",
    roomName: "",
    selectedDate: "",
    selectedTime: "",
    selectedPrice: "",
    selectedFoodOptions: [],
  }),
  actions: {
    setBookingDetails(details) {
      this.roomID = details.roomID;
      this.roomName = details.roomName;
      this.selectedDate = details.selectedDate;
      this.selectedTime = details.selectedTime;
      this.selectedPrice = details.selectedPrice;
      this.selectedFoodOptions = details.selectedFoodOptions;
    }
  }
});
