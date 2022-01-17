# from locust import Locust, TaskSet, task
from locust import HttpUser, TaskSet, task

class DummyTaskSet(TaskSet):
    @task()
    def dummy_pass(self):
        pass

class Dummy(HttpUser):
    task_set = DummyTaskSet
