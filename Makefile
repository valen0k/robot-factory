OBJS = updateStorageCost updateNumberRobots

all:
	go build cmd/updateNumberRobots.go
	go build cmd/updateStorageCost.go
	go run cmd/robotService.go

clean:

fclean:	clean
	rm -f $(OBJS)

re:	fclean all

.PHONY:	all clean fclean re