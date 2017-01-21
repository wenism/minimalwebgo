FROM scratch
ADD minimalweb /
EXPOSE 80
CMD ["/minimalweb"]