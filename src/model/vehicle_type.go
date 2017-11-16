package model

/**
 * Тип техники.
 */
type VehicleType uint8

const Vehicle_None VehicleType = 255

const (
	/**
	 * Бронированная ремонтно-эвакуационная машина. Наземный юнит.
	 * Постепенно восстанавливает прочность находящейся поблизости неподвижной техники.
	 */
	Vehicle_Arrv VehicleType = iota

	/**
	 * Истребитель. Воздушный юнит. Крайне эффективен против другой воздушной техники. Не может атаковать наземные цели.
	 */
	Vehicle_Fighter

	/**
	 * Ударный вертолёт. Воздушный юнит. Может атаковать как воздушные, так и наземные цели.
	 */
	Vehicle_Helicopter

	/**
	 * Боевая машина пехоты. Наземный юнит. Может атаковать как воздушные, так и наземные цели.
	 */
	Vehicle_Ifv

	/**
	 * Танк. Наземный юнит. Крайне эффективен против другой наземной техники. Также может атаковать воздушные цели.
	 */
	Vehicle_Tank
)
