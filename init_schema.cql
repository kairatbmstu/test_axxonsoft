CREATE KEYSPACE test_axxonsoft
    WITH replication = {'class': 'NetworkTopologyStrategy', 'replication_factor' : 3};
