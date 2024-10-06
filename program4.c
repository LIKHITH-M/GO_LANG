//  Design, Develop and implement menu driven program to simulate processing of 
// batch jobs by a computer system. The scheduling of these jobs should be handled 
// using a priority queue.

// Note: The Program should allow users to add or remove items from the queue and it should 
// also display current status i.e. the total number of items in the queue.
// #include <stdio.h>
#include <stdlib.h>
#include<stdio.h>
#define MAX 5

int front = -1, rear = -1;

typedef struct process {
    int pid;
    int pr;
    int bt;
} job;

job pjob[MAX];

// Function to insert a job into the priority queue
void insert() {
    int pid, pr, bt;
    
    if (rear == MAX - 1) {
        printf("Queue Overflow\n");
        return;
    }
    
    printf("Enter PID, PR and BT: ");
    scanf("%d %d %d", &pid, &pr, &bt);
    
    if (rear == -1) {
        front = 0;
    }
    
    rear++;
    pjob[rear].pid = pid;
    pjob[rear].pr = pr;
    pjob[rear].bt = bt;
}

// Function to delete the job with the highest priority
void delete() {
    int i, pos = 0, max = -1;
    
    if (front == -1) {
        printf("Queue Underflow\n");
        return;
    }
    
    if (front == rear) {
        front = -1;
        rear = -1;
        return;
    }
    
    // Find the position of the job with the highest priority
    for (i = front; i <= rear; i++) {
        if (pjob[i].pr > max) {
            max = pjob[i].pr;
            pos = i;
        }
    }
    
    // Shift elements to remove the job with highest priority
    for (i = pos; i < rear; i++) {
        pjob[i] = pjob[i + 1];
    }
    
    rear--;
}

// Function to display all jobs in the queue
void display() {
    if (front == -1) {
        printf("Queue is Empty\n");
        return;
    }
    
    printf("PID\tPR\tBT\n");
    for (int i = front; i <= rear; i++) {
        printf("%d\t%d\t%d\n", pjob[i].pid, pjob[i].pr, pjob[i].bt);
    }
}

// Main function to drive the menu-driven program
int main() {
    int ch;
    
    while (1) {
        printf("\n1. Insert\t2. Display\t3. Delete\t4. Exit\n");
        printf("Enter your choice: ");
        scanf("%d", &ch);
        
        switch (ch) {
            case 1: 
                insert();
                break;
            case 2: 
                display();
                break;
            case 3: 
                delete();
                break;
            case 4: 
                exit(0);
                break;
            default: 
                printf("Invalid choice\n");
                break;
        }
    }
    
    return 0;
}
