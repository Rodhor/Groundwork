<template>
  <v-container
    max-width="1500"
    class="d-flex flex-column align-center justify-center"
  >
    <div>
      <div class="text-center mb-5 pt-5 text-h4">
        <h1>Participants Overview</h1>
      </div>
    </div>
    <v-card floating outlined class="mt-10 w-100">
      <v-card-title class="d-flex align-center pe-2">
        <v-icon icon="mdi-account-multiple"></v-icon> &nbsp; Find a Participant

        <v-spacer></v-spacer>

        <v-text-field
          v-model="search"
          density="compact"
          label="Search"
          prepend-inner-icon="mdi-magnify"
          variant="solo-filled"
          flat
          hide-details
          single-line
        ></v-text-field>
      </v-card-title>

      <v-divider></v-divider>
      <v-data-table
        v-model:search="search"
        :filter-keys="['name']"
        :headers="headers"
        :items="items"
        hover
        no-data-text="No Participants currently"
        striped
        width="100%"
      >
        <template v-slot:item.status="{ item }">
          <div>
            <v-chip
              :color="statusChip(item).color"
              :text="statusChip(item).label"
              class="text-uppercase"
              size="small"
              label
            ></v-chip>
          </div>
        </template>
      </v-data-table>
    </v-card>
  </v-container>
</template>

<script setup>
import { ref } from "vue";

const search = ref("");

const headers = [
  { title: "Name", value: "name" },
  { title: "Email", value: "email" },
  { title: "Start Date", value: "startDate" },
  { title: "End Date", value: "endDate" },
  { title: "Status", value: "status" },
];

const items = [
  // Found work (Green)
  {
    name: "Alice Employed",
    email: "alice@work.com",
    startDate: "2025-10-01",
    endDate: "2026-01-01",
    foundWork: "true",
  },
  // Currently participating (Blue)
  {
    name: "Bob Current",
    email: "bob@now.com",
    startDate: "2025-11-01",
    endDate: "2026-02-01",
    foundWork: "false",
  },
  // In the past, not found work (Red)
  {
    name: "Charlie Ended",
    email: "charlie@past.com",
    startDate: "2025-06-01",
    endDate: "2025-09-01",
    foundWork: "false",
  },
  // In the future (Yellow)
  {
    name: "Dana Future",
    email: "dana@soon.com",
    startDate: "2026-05-01",
    endDate: "2026-08-01",
    foundWork: "false",
  },
];

// Date status logic
function statusChip(item) {
  const today = new Date();
  today.setHours(0, 0, 0, 0);
  const start = new Date(item.startDate);
  const end = new Date(item.endDate);
  start.setHours(0, 0, 0, 0);
  end.setHours(0, 0, 0, 0);

  if (item.foundWork === "true") {
    return { color: "green", label: "found work" };
  } else if (today < start) {
    return { color: "yellow", label: "future" };
  } else if (today > end) {
    return { color: "red", label: "ended" };
  } else if (today >= start && today <= end) {
    return { color: "blue", label: "participating" };
  }
  return { color: "grey", label: "unknown" };
}
</script>
