<template>
  <div>
    <v-container
      fluid
      max-width="1500"
      class="d-flex flex-column align-center justify-center"
    >
      <div class="text-center mb-5 pt-5 text-h4">
        <h1>Tasks Overview</h1>
      </div>
      <v-spacer></v-spacer>
      <v-card floating outlined class="mt-10 w-100 pa-4">
        <div v-if="tasks.length > 0" class="w-100">
          <v-table hover striped="even" class="w-100">
            <thead>
              <tr>
                <th class="text-center">Task</th>
                <th class="text-center">Description</th>
                <th class="text-center">Status</th>
                <th class="text-center">Due Date</th>
                <th class="text-center">Actions</th>
              </tr>
            </thead>
            <tbody>
              <tr v-for="item in tasks" :key="item.id">
                <td class="text-center">{{ item.name }}</td>
                <td class="text-center">{{ item.description }}</td>
                <td class="text-center">
                  <v-chip
                    v-if="
                      new Date() > new Date(item.dueDate) &&
                      item.status !== 'Completed'
                    "
                    color="red"
                    dark
                    @click="
                      item.status =
                        item.status === 'Completed' ? 'Pending' : 'Completed'
                    "
                  >
                    Overdue
                  </v-chip>
                  <v-chip
                    v-else
                    :color="item.status === 'Completed' ? 'green' : 'orange'"
                    dark
                    @click="
                      item.status =
                        item.status === 'Completed' ? 'Pending' : 'Completed'
                    "
                  >
                    {{ item.status }}
                  </v-chip>
                </td>
                <td class="text-center">{{ item.dueDate }}</td>
                <td class="text-center">
                  <v-btn class="mr-2" color="primary" @click="editTask(item)"
                    >Edit</v-btn
                  >
                  <v-btn class="mr-2" color="error" @click="deleteTask(item)"
                    >Delete</v-btn
                  >
                </td>
              </tr>
            </tbody>
          </v-table>
        </div>
        <div v-else class="flex flex-column align-center text-center">
          <div class="mb-4">
            <v-icon size="64" color="grey"
              >mdi-checkbox-marked-circle-outline</v-icon
            >
          </div>
          <h2>No tasks available.</h2>
        </div>
      </v-card>
    </v-container>
  </div>
</template>

<script>
export default {
  data() {
    return {
      tasks: [
        {
          id: 1,
          name: "Task 1",
          description: "Description for Task 1",
          status: "Pending",
          dueDate: "2024-06-10",
        },
        {
          id: 2,
          name: "Task 2",
          description: "Description for Task 2",
          status: "Completed",
          dueDate: "2026-06-12",
        },
        {
          id: 3,
          name: "Task 3",
          description: "Description for Task 3",
          status: "Pending",
          dueDate: "2024-06-15",
        },
      ],
    };
  },
  methods: {
    editTask(task) {
      // Implement edit functionality
    },
    deleteTask(task) {
      this.tasks = this.tasks.filter((t) => t.id !== task.id);
    },
  },
};
</script>

<style scoped></style>
